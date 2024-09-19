package validate

import (
	"fmt"
	protosNearUsers "github.com/martbul/near_users/protos/near_users"
)

// validateLocation ensures that the coordinates are within valid ranges.
func ValidateLocation(loc *protosNearUsers.UserTokenAndLocation) error {
	if loc.Latitude < -90 || loc.Latitude > 90 {
		return fmt.Errorf("latitude must be between -90 and 90")
	}
	if loc.Longitude < -180 || loc.Longitude > 180 {
		return fmt.Errorf("longitude must be between -180 and 180")
	}
	return nil
}
