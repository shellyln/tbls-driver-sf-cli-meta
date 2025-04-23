package driver

import "reflect"

func WalkAndApply(rootName string, value any, trans func(ancestors []string, parent string, ty reflect.Type, value any) any) any {
	return walkAndApplyInternal(make([]string, 0), rootName, reflect.TypeOf(value), reflect.ValueOf(value), trans)
}

func walkAndApplyInternal(
	ancestors []string, parent string, ty reflect.Type, value reflect.Value,
	trans func(ancestors []string, parent string, ty reflect.Type, value any) any) reflect.Value {

	if ty.Kind() == reflect.Ptr {
		derefTy := ty.Elem()
		switch derefTy.Kind() {
		case reflect.Struct, reflect.Slice, reflect.Array, reflect.Map:
			// derefer
			if !value.IsNil() {
				walkAndApplyInternal(ancestors, parent, derefTy, value.Elem(), trans)
			}
			return value
		}
	}

	switch ty.Kind() {
	case reflect.Slice, reflect.Map:
		if value.IsNil() {
			return value
		}
	}

	switch ty.Kind() {
	case reflect.Struct:
		if len(parent) > 0 {
			ancestors = append(ancestors, parent)
		}
		for i := range ty.NumField() {
			field := ty.Field(i)
			if field.PkgPath != "" {
				// private field
				continue
			}
			rv := value.Field(i)
			if !(rv.IsValid() && rv.CanSet()) {
				continue
			}
			rt := field.Type
			if value.Interface() != nil {
				rt = rv.Type()
			}
			newValue := walkAndApplyInternal(ancestors, field.Name, rt, rv, trans)
			rv.Set(newValue)
		}
		return value
	case reflect.Slice, reflect.Array:
		srt := ty.Elem()
		for i := range value.Len() {
			rv := value.Index(i)
			if !(rv.IsValid() && rv.CanSet()) {
				continue
			}
			rt := srt
			if value.Interface() != nil {
				rt = rv.Type()
			}
			newValue := walkAndApplyInternal(ancestors, parent, rt, rv, trans)
			rv.Set(newValue)
		}
		return value
	case reflect.Map:
		mrt := ty.Elem()
		iter := value.MapRange()
		for iter.Next() {
			elemKey := iter.Key().Interface()
			rv := iter.Value()
			rt := mrt
			if value.Interface() != nil {
				rt = rv.Type()
			}
			newValue := walkAndApplyInternal(ancestors, parent, rt, rv, trans)
			value.SetMapIndex(reflect.ValueOf(elemKey), newValue)
		}
		return value
	case reflect.Ptr:
		prt := ty.Elem()
		rv := value.Elem()
		if !(rv.IsValid() && rv.CanSet()) {
			break
		}
		rt := prt
		if value.Interface() != nil {
			rt = rv.Type()
		}
		newValue := walkAndApplyInternal(ancestors, parent, rt, rv, trans)
		rv.Set(newValue)
		return value
	default:
		rt := ty
		if value.Interface() != nil {
			rt = value.Type()
		}
		newValue := trans(ancestors, parent, rt, value.Interface())
		if newValue == nil {
			tmpptr := &newValue
			tmprv := reflect.ValueOf(tmpptr)
			return tmprv.Elem()
		} else {
			return reflect.ValueOf(newValue)
		}
	}

	return reflect.ValueOf(nil)
}
