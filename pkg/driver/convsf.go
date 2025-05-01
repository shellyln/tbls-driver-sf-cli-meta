package driver

func (s *SfSharingBaseCriteriaRule) getCriteria() string {
	desc := s.BooleanFilter

	for _, v := range s.CriteriaItems {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += v.Field + " " + v.Operation + " " + v.Value + v.ValueField
	}

	return "{" + desc + "}"
}

func (s *SfSharingCriteriaRule) ToDescription() string {
	desc := "[" + s.AccessLevel + "]"

	if s.IncludeRecordsOwnedByAll {
		desc += "[IncludeRecordsOwnedByAll]"
	}
	desc += "; From " + s.getCriteria() + "; To " + s.SharedTo.ToDescription()

	return desc
}

func (s *SfSharingGuestRules) ToDescription() string {
	desc := "[" + s.AccessLevel + "]"

	if s.IncludeHVUOwnedRecords {
		desc += "[IncludeHVUOwnedRecords]"
	}
	desc += "; From " + s.getCriteria() + "; To " + s.SharedTo.ToDescription()

	return desc
}

func (s *SfSharingOwnerRules) ToDescription() string {
	return "[" + s.AccessLevel + "]; From " + s.SharedFrom.ToDescription() + "; To " + s.SharedTo.ToDescription()
}

func (s *SfSharedTo) ToDescription() string {
	desc := ""

	if s.AllCustomerPortalUsers != nil {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "AllCustomerPortalUsers"
	}
	if s.AllInternalUsers != nil {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "AllInternalUsers"
	}
	if s.AllPartnerUsers != nil {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "AllPartnerUsers"
	}
	if s.ChannelProgramGroup != nil {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "ChannelProgramGroup"
	}

	if len(s.Group) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "Group{"
		for i, v := range s.Group {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.GuestUser) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "GuestUser{"
		for i, v := range s.GuestUser {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.ManagerSubordinates) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "ManagerSubordinates{"
		for i, v := range s.ManagerSubordinates {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.Managers) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "Managers{"
		for i, v := range s.Managers {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.PortalRole) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "PortalRole{"
		for i, v := range s.PortalRole {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.PortalRoleAndSubordinates) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "PortalRoleAndSubordinates{"
		for i, v := range s.PortalRoleAndSubordinates {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.Role) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "Role{"
		for i, v := range s.Role {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.RoleAndSubordinates) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "RoleAndSubordinates{"
		for i, v := range s.RoleAndSubordinates {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.RoleAndSubordinatesInternal) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "RoleAndSubordinatesInternal{"
		for i, v := range s.RoleAndSubordinatesInternal {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.Territory) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "Territory{"
		for i, v := range s.Territory {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.TerritoryAndSubordinates) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "TerritoryAndSubordinates{"
		for i, v := range s.TerritoryAndSubordinates {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}
	if len(s.Queue) > 0 {
		if len(desc) > 0 {
			desc += ", "
		}
		desc += "Queue{"
		for i, v := range s.Queue {
			if i != 0 {
				desc += ", "
			}
			desc += v
		}
		desc += "}"
	}

	return desc
}
