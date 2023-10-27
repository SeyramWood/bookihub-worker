// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/SeyramWood/ent/bookibususer"
	"github.com/SeyramWood/ent/booking"
	"github.com/SeyramWood/ent/company"
	"github.com/SeyramWood/ent/companyuser"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/customercontact"
	"github.com/SeyramWood/ent/customerluggage"
	"github.com/SeyramWood/ent/notification"
	"github.com/SeyramWood/ent/passenger"
	"github.com/SeyramWood/ent/route"
	"github.com/SeyramWood/ent/routestop"
	"github.com/SeyramWood/ent/schema"
	"github.com/SeyramWood/ent/trip"
	"github.com/SeyramWood/ent/user"
	"github.com/SeyramWood/ent/vehicle"
	"github.com/SeyramWood/ent/vehicleimage"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bookibususerMixin := schema.BookibusUser{}.Mixin()
	bookibususerMixinFields0 := bookibususerMixin[0].Fields()
	_ = bookibususerMixinFields0
	bookibususerFields := schema.BookibusUser{}.Fields()
	_ = bookibususerFields
	// bookibususerDescCreatedAt is the schema descriptor for created_at field.
	bookibususerDescCreatedAt := bookibususerMixinFields0[0].Descriptor()
	// bookibususer.DefaultCreatedAt holds the default value on creation for the created_at field.
	bookibususer.DefaultCreatedAt = bookibususerDescCreatedAt.Default.(func() time.Time)
	// bookibususerDescUpdatedAt is the schema descriptor for updated_at field.
	bookibususerDescUpdatedAt := bookibususerMixinFields0[1].Descriptor()
	// bookibususer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	bookibususer.DefaultUpdatedAt = bookibususerDescUpdatedAt.Default.(func() time.Time)
	// bookibususer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	bookibususer.UpdateDefaultUpdatedAt = bookibususerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// bookibususerDescLastName is the schema descriptor for last_name field.
	bookibususerDescLastName := bookibususerFields[0].Descriptor()
	// bookibususer.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	bookibususer.LastNameValidator = bookibususerDescLastName.Validators[0].(func(string) error)
	// bookibususerDescOtherName is the schema descriptor for other_name field.
	bookibususerDescOtherName := bookibususerFields[1].Descriptor()
	// bookibususer.OtherNameValidator is a validator for the "other_name" field. It is called by the builders before save.
	bookibususer.OtherNameValidator = bookibususerDescOtherName.Validators[0].(func(string) error)
	// bookibususerDescPhone is the schema descriptor for phone field.
	bookibususerDescPhone := bookibususerFields[2].Descriptor()
	// bookibususer.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	bookibususer.PhoneValidator = bookibususerDescPhone.Validators[0].(func(string) error)
	bookingMixin := schema.Booking{}.Mixin()
	bookingMixinFields0 := bookingMixin[0].Fields()
	_ = bookingMixinFields0
	bookingFields := schema.Booking{}.Fields()
	_ = bookingFields
	// bookingDescCreatedAt is the schema descriptor for created_at field.
	bookingDescCreatedAt := bookingMixinFields0[0].Descriptor()
	// booking.DefaultCreatedAt holds the default value on creation for the created_at field.
	booking.DefaultCreatedAt = bookingDescCreatedAt.Default.(func() time.Time)
	// bookingDescUpdatedAt is the schema descriptor for updated_at field.
	bookingDescUpdatedAt := bookingMixinFields0[1].Descriptor()
	// booking.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	booking.DefaultUpdatedAt = bookingDescUpdatedAt.Default.(func() time.Time)
	// booking.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	booking.UpdateDefaultUpdatedAt = bookingDescUpdatedAt.UpdateDefault.(func() time.Time)
	// bookingDescBookingNumber is the schema descriptor for booking_number field.
	bookingDescBookingNumber := bookingFields[0].Descriptor()
	// booking.BookingNumberValidator is a validator for the "booking_number" field. It is called by the builders before save.
	booking.BookingNumberValidator = bookingDescBookingNumber.Validators[0].(func(string) error)
	// bookingDescBoardingPoint is the schema descriptor for boarding_point field.
	bookingDescBoardingPoint := bookingFields[1].Descriptor()
	// booking.BoardingPointValidator is a validator for the "boarding_point" field. It is called by the builders before save.
	booking.BoardingPointValidator = bookingDescBoardingPoint.Validators[0].(func(string) error)
	// bookingDescVat is the schema descriptor for vat field.
	bookingDescVat := bookingFields[2].Descriptor()
	// booking.DefaultVat holds the default value on creation for the vat field.
	booking.DefaultVat = bookingDescVat.Default.(float64)
	// bookingDescSmsFee is the schema descriptor for sms_fee field.
	bookingDescSmsFee := bookingFields[3].Descriptor()
	// booking.DefaultSmsFee holds the default value on creation for the sms_fee field.
	booking.DefaultSmsFee = bookingDescSmsFee.Default.(float64)
	// bookingDescAmount is the schema descriptor for amount field.
	bookingDescAmount := bookingFields[4].Descriptor()
	// booking.DefaultAmount holds the default value on creation for the amount field.
	booking.DefaultAmount = bookingDescAmount.Default.(float64)
	// bookingDescSmsNotification is the schema descriptor for sms_notification field.
	bookingDescSmsNotification := bookingFields[8].Descriptor()
	// booking.DefaultSmsNotification holds the default value on creation for the sms_notification field.
	booking.DefaultSmsNotification = bookingDescSmsNotification.Default.(bool)
	companyMixin := schema.Company{}.Mixin()
	companyMixinFields0 := companyMixin[0].Fields()
	_ = companyMixinFields0
	companyFields := schema.Company{}.Fields()
	_ = companyFields
	// companyDescCreatedAt is the schema descriptor for created_at field.
	companyDescCreatedAt := companyMixinFields0[0].Descriptor()
	// company.DefaultCreatedAt holds the default value on creation for the created_at field.
	company.DefaultCreatedAt = companyDescCreatedAt.Default.(func() time.Time)
	// companyDescUpdatedAt is the schema descriptor for updated_at field.
	companyDescUpdatedAt := companyMixinFields0[1].Descriptor()
	// company.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	company.DefaultUpdatedAt = companyDescUpdatedAt.Default.(func() time.Time)
	// company.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	company.UpdateDefaultUpdatedAt = companyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// companyDescName is the schema descriptor for name field.
	companyDescName := companyFields[0].Descriptor()
	// company.NameValidator is a validator for the "name" field. It is called by the builders before save.
	company.NameValidator = companyDescName.Validators[0].(func(string) error)
	// companyDescPhone is the schema descriptor for phone field.
	companyDescPhone := companyFields[1].Descriptor()
	// company.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	company.PhoneValidator = companyDescPhone.Validators[0].(func(string) error)
	// companyDescEmail is the schema descriptor for email field.
	companyDescEmail := companyFields[3].Descriptor()
	// company.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	company.EmailValidator = companyDescEmail.Validators[0].(func(string) error)
	companyuserMixin := schema.CompanyUser{}.Mixin()
	companyuserMixinFields0 := companyuserMixin[0].Fields()
	_ = companyuserMixinFields0
	companyuserFields := schema.CompanyUser{}.Fields()
	_ = companyuserFields
	// companyuserDescCreatedAt is the schema descriptor for created_at field.
	companyuserDescCreatedAt := companyuserMixinFields0[0].Descriptor()
	// companyuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	companyuser.DefaultCreatedAt = companyuserDescCreatedAt.Default.(func() time.Time)
	// companyuserDescUpdatedAt is the schema descriptor for updated_at field.
	companyuserDescUpdatedAt := companyuserMixinFields0[1].Descriptor()
	// companyuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	companyuser.DefaultUpdatedAt = companyuserDescUpdatedAt.Default.(func() time.Time)
	// companyuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	companyuser.UpdateDefaultUpdatedAt = companyuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	customerMixin := schema.Customer{}.Mixin()
	customerMixinFields0 := customerMixin[0].Fields()
	_ = customerMixinFields0
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescCreatedAt is the schema descriptor for created_at field.
	customerDescCreatedAt := customerMixinFields0[0].Descriptor()
	// customer.DefaultCreatedAt holds the default value on creation for the created_at field.
	customer.DefaultCreatedAt = customerDescCreatedAt.Default.(func() time.Time)
	// customerDescUpdatedAt is the schema descriptor for updated_at field.
	customerDescUpdatedAt := customerMixinFields0[1].Descriptor()
	// customer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	customer.DefaultUpdatedAt = customerDescUpdatedAt.Default.(func() time.Time)
	// customer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	customer.UpdateDefaultUpdatedAt = customerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// customerDescLastName is the schema descriptor for last_name field.
	customerDescLastName := customerFields[0].Descriptor()
	// customer.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	customer.LastNameValidator = customerDescLastName.Validators[0].(func(string) error)
	// customerDescOtherName is the schema descriptor for other_name field.
	customerDescOtherName := customerFields[1].Descriptor()
	// customer.OtherNameValidator is a validator for the "other_name" field. It is called by the builders before save.
	customer.OtherNameValidator = customerDescOtherName.Validators[0].(func(string) error)
	// customerDescPhone is the schema descriptor for phone field.
	customerDescPhone := customerFields[2].Descriptor()
	// customer.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	customer.PhoneValidator = customerDescPhone.Validators[0].(func(string) error)
	customercontactMixin := schema.CustomerContact{}.Mixin()
	customercontactMixinFields0 := customercontactMixin[0].Fields()
	_ = customercontactMixinFields0
	customercontactFields := schema.CustomerContact{}.Fields()
	_ = customercontactFields
	// customercontactDescCreatedAt is the schema descriptor for created_at field.
	customercontactDescCreatedAt := customercontactMixinFields0[0].Descriptor()
	// customercontact.DefaultCreatedAt holds the default value on creation for the created_at field.
	customercontact.DefaultCreatedAt = customercontactDescCreatedAt.Default.(func() time.Time)
	// customercontactDescUpdatedAt is the schema descriptor for updated_at field.
	customercontactDescUpdatedAt := customercontactMixinFields0[1].Descriptor()
	// customercontact.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	customercontact.DefaultUpdatedAt = customercontactDescUpdatedAt.Default.(func() time.Time)
	// customercontact.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	customercontact.UpdateDefaultUpdatedAt = customercontactDescUpdatedAt.UpdateDefault.(func() time.Time)
	// customercontactDescFullName is the schema descriptor for full_name field.
	customercontactDescFullName := customercontactFields[0].Descriptor()
	// customercontact.FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	customercontact.FullNameValidator = customercontactDescFullName.Validators[0].(func(string) error)
	// customercontactDescEmail is the schema descriptor for email field.
	customercontactDescEmail := customercontactFields[1].Descriptor()
	// customercontact.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	customercontact.EmailValidator = customercontactDescEmail.Validators[0].(func(string) error)
	// customercontactDescPhone is the schema descriptor for phone field.
	customercontactDescPhone := customercontactFields[2].Descriptor()
	// customercontact.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	customercontact.PhoneValidator = customercontactDescPhone.Validators[0].(func(string) error)
	customerluggageMixin := schema.CustomerLuggage{}.Mixin()
	customerluggageMixinFields0 := customerluggageMixin[0].Fields()
	_ = customerluggageMixinFields0
	customerluggageFields := schema.CustomerLuggage{}.Fields()
	_ = customerluggageFields
	// customerluggageDescCreatedAt is the schema descriptor for created_at field.
	customerluggageDescCreatedAt := customerluggageMixinFields0[0].Descriptor()
	// customerluggage.DefaultCreatedAt holds the default value on creation for the created_at field.
	customerluggage.DefaultCreatedAt = customerluggageDescCreatedAt.Default.(func() time.Time)
	// customerluggageDescUpdatedAt is the schema descriptor for updated_at field.
	customerluggageDescUpdatedAt := customerluggageMixinFields0[1].Descriptor()
	// customerluggage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	customerluggage.DefaultUpdatedAt = customerluggageDescUpdatedAt.Default.(func() time.Time)
	// customerluggage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	customerluggage.UpdateDefaultUpdatedAt = customerluggageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// customerluggageDescQuantity is the schema descriptor for quantity field.
	customerluggageDescQuantity := customerluggageFields[1].Descriptor()
	// customerluggage.DefaultQuantity holds the default value on creation for the quantity field.
	customerluggage.DefaultQuantity = customerluggageDescQuantity.Default.(int)
	// customerluggageDescAmount is the schema descriptor for amount field.
	customerluggageDescAmount := customerluggageFields[2].Descriptor()
	// customerluggage.DefaultAmount holds the default value on creation for the amount field.
	customerluggage.DefaultAmount = customerluggageDescAmount.Default.(float64)
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescEvent is the schema descriptor for event field.
	notificationDescEvent := notificationFields[0].Descriptor()
	// notification.EventValidator is a validator for the "event" field. It is called by the builders before save.
	notification.EventValidator = notificationDescEvent.Validators[0].(func(string) error)
	// notificationDescActivity is the schema descriptor for activity field.
	notificationDescActivity := notificationFields[1].Descriptor()
	// notification.ActivityValidator is a validator for the "activity" field. It is called by the builders before save.
	notification.ActivityValidator = notificationDescActivity.Validators[0].(func(string) error)
	// notificationDescDescription is the schema descriptor for description field.
	notificationDescDescription := notificationFields[2].Descriptor()
	// notification.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	notification.DescriptionValidator = notificationDescDescription.Validators[0].(func(string) error)
	// notificationDescSubjectType is the schema descriptor for subject_type field.
	notificationDescSubjectType := notificationFields[3].Descriptor()
	// notification.SubjectTypeValidator is a validator for the "subject_type" field. It is called by the builders before save.
	notification.SubjectTypeValidator = notificationDescSubjectType.Validators[0].(func(string) error)
	// notificationDescCreatorType is the schema descriptor for creator_type field.
	notificationDescCreatorType := notificationFields[5].Descriptor()
	// notification.CreatorTypeValidator is a validator for the "creator_type" field. It is called by the builders before save.
	notification.CreatorTypeValidator = notificationDescCreatorType.Validators[0].(func(string) error)
	passengerMixin := schema.Passenger{}.Mixin()
	passengerMixinFields0 := passengerMixin[0].Fields()
	_ = passengerMixinFields0
	passengerFields := schema.Passenger{}.Fields()
	_ = passengerFields
	// passengerDescCreatedAt is the schema descriptor for created_at field.
	passengerDescCreatedAt := passengerMixinFields0[0].Descriptor()
	// passenger.DefaultCreatedAt holds the default value on creation for the created_at field.
	passenger.DefaultCreatedAt = passengerDescCreatedAt.Default.(func() time.Time)
	// passengerDescUpdatedAt is the schema descriptor for updated_at field.
	passengerDescUpdatedAt := passengerMixinFields0[1].Descriptor()
	// passenger.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	passenger.DefaultUpdatedAt = passengerDescUpdatedAt.Default.(func() time.Time)
	// passenger.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	passenger.UpdateDefaultUpdatedAt = passengerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// passengerDescFullName is the schema descriptor for full_name field.
	passengerDescFullName := passengerFields[0].Descriptor()
	// passenger.FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	passenger.FullNameValidator = passengerDescFullName.Validators[0].(func(string) error)
	// passengerDescAmount is the schema descriptor for amount field.
	passengerDescAmount := passengerFields[1].Descriptor()
	// passenger.DefaultAmount holds the default value on creation for the amount field.
	passenger.DefaultAmount = passengerDescAmount.Default.(float64)
	routeMixin := schema.Route{}.Mixin()
	routeMixinFields0 := routeMixin[0].Fields()
	_ = routeMixinFields0
	routeFields := schema.Route{}.Fields()
	_ = routeFields
	// routeDescCreatedAt is the schema descriptor for created_at field.
	routeDescCreatedAt := routeMixinFields0[0].Descriptor()
	// route.DefaultCreatedAt holds the default value on creation for the created_at field.
	route.DefaultCreatedAt = routeDescCreatedAt.Default.(func() time.Time)
	// routeDescUpdatedAt is the schema descriptor for updated_at field.
	routeDescUpdatedAt := routeMixinFields0[1].Descriptor()
	// route.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	route.DefaultUpdatedAt = routeDescUpdatedAt.Default.(func() time.Time)
	// route.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	route.UpdateDefaultUpdatedAt = routeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// routeDescFromLocation is the schema descriptor for from_location field.
	routeDescFromLocation := routeFields[0].Descriptor()
	// route.FromLocationValidator is a validator for the "from_location" field. It is called by the builders before save.
	route.FromLocationValidator = routeDescFromLocation.Validators[0].(func(string) error)
	// routeDescToLocation is the schema descriptor for to_location field.
	routeDescToLocation := routeFields[1].Descriptor()
	// route.ToLocationValidator is a validator for the "to_location" field. It is called by the builders before save.
	route.ToLocationValidator = routeDescToLocation.Validators[0].(func(string) error)
	// routeDescRate is the schema descriptor for rate field.
	routeDescRate := routeFields[6].Descriptor()
	// route.DefaultRate holds the default value on creation for the rate field.
	route.DefaultRate = routeDescRate.Default.(float64)
	// routeDescDiscount is the schema descriptor for discount field.
	routeDescDiscount := routeFields[7].Descriptor()
	// route.DefaultDiscount holds the default value on creation for the discount field.
	route.DefaultDiscount = routeDescDiscount.Default.(float32)
	// routeDescPopularity is the schema descriptor for popularity field.
	routeDescPopularity := routeFields[8].Descriptor()
	// route.DefaultPopularity holds the default value on creation for the popularity field.
	route.DefaultPopularity = routeDescPopularity.Default.(int)
	routestopMixin := schema.RouteStop{}.Mixin()
	routestopMixinFields0 := routestopMixin[0].Fields()
	_ = routestopMixinFields0
	routestopFields := schema.RouteStop{}.Fields()
	_ = routestopFields
	// routestopDescCreatedAt is the schema descriptor for created_at field.
	routestopDescCreatedAt := routestopMixinFields0[0].Descriptor()
	// routestop.DefaultCreatedAt holds the default value on creation for the created_at field.
	routestop.DefaultCreatedAt = routestopDescCreatedAt.Default.(func() time.Time)
	// routestopDescUpdatedAt is the schema descriptor for updated_at field.
	routestopDescUpdatedAt := routestopMixinFields0[1].Descriptor()
	// routestop.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	routestop.DefaultUpdatedAt = routestopDescUpdatedAt.Default.(func() time.Time)
	// routestop.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	routestop.UpdateDefaultUpdatedAt = routestopDescUpdatedAt.UpdateDefault.(func() time.Time)
	tripMixin := schema.Trip{}.Mixin()
	tripMixinFields0 := tripMixin[0].Fields()
	_ = tripMixinFields0
	tripFields := schema.Trip{}.Fields()
	_ = tripFields
	// tripDescCreatedAt is the schema descriptor for created_at field.
	tripDescCreatedAt := tripMixinFields0[0].Descriptor()
	// trip.DefaultCreatedAt holds the default value on creation for the created_at field.
	trip.DefaultCreatedAt = tripDescCreatedAt.Default.(func() time.Time)
	// tripDescUpdatedAt is the schema descriptor for updated_at field.
	tripDescUpdatedAt := tripMixinFields0[1].Descriptor()
	// trip.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	trip.DefaultUpdatedAt = tripDescUpdatedAt.Default.(func() time.Time)
	// trip.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	trip.UpdateDefaultUpdatedAt = tripDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tripDescExteriorInspected is the schema descriptor for exterior_inspected field.
	tripDescExteriorInspected := tripFields[4].Descriptor()
	// trip.DefaultExteriorInspected holds the default value on creation for the exterior_inspected field.
	trip.DefaultExteriorInspected = tripDescExteriorInspected.Default.(bool)
	// tripDescInteriorInspected is the schema descriptor for interior_inspected field.
	tripDescInteriorInspected := tripFields[5].Descriptor()
	// trip.DefaultInteriorInspected holds the default value on creation for the interior_inspected field.
	trip.DefaultInteriorInspected = tripDescInteriorInspected.Default.(bool)
	// tripDescEngineCompartmentInspected is the schema descriptor for engine_compartment_inspected field.
	tripDescEngineCompartmentInspected := tripFields[6].Descriptor()
	// trip.DefaultEngineCompartmentInspected holds the default value on creation for the engine_compartment_inspected field.
	trip.DefaultEngineCompartmentInspected = tripDescEngineCompartmentInspected.Default.(bool)
	// tripDescBrakeAndSteeringInspected is the schema descriptor for brake_and_steering_inspected field.
	tripDescBrakeAndSteeringInspected := tripFields[7].Descriptor()
	// trip.DefaultBrakeAndSteeringInspected holds the default value on creation for the brake_and_steering_inspected field.
	trip.DefaultBrakeAndSteeringInspected = tripDescBrakeAndSteeringInspected.Default.(bool)
	// tripDescEmergencyEquipmentInspected is the schema descriptor for emergency_equipment_inspected field.
	tripDescEmergencyEquipmentInspected := tripFields[8].Descriptor()
	// trip.DefaultEmergencyEquipmentInspected holds the default value on creation for the emergency_equipment_inspected field.
	trip.DefaultEmergencyEquipmentInspected = tripDescEmergencyEquipmentInspected.Default.(bool)
	// tripDescFuelAndFluidsInspected is the schema descriptor for fuel_and_fluids_inspected field.
	tripDescFuelAndFluidsInspected := tripFields[9].Descriptor()
	// trip.DefaultFuelAndFluidsInspected holds the default value on creation for the fuel_and_fluids_inspected field.
	trip.DefaultFuelAndFluidsInspected = tripDescFuelAndFluidsInspected.Default.(bool)
	// tripDescScheduled is the schema descriptor for scheduled field.
	tripDescScheduled := tripFields[10].Descriptor()
	// trip.DefaultScheduled holds the default value on creation for the scheduled field.
	trip.DefaultScheduled = tripDescScheduled.Default.(bool)
	// tripDescSeatLeft is the schema descriptor for seat_left field.
	tripDescSeatLeft := tripFields[11].Descriptor()
	// trip.DefaultSeatLeft holds the default value on creation for the seat_left field.
	trip.DefaultSeatLeft = tripDescSeatLeft.Default.(int)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func([]byte) error)
	vehicleMixin := schema.Vehicle{}.Mixin()
	vehicleMixinFields0 := vehicleMixin[0].Fields()
	_ = vehicleMixinFields0
	vehicleFields := schema.Vehicle{}.Fields()
	_ = vehicleFields
	// vehicleDescCreatedAt is the schema descriptor for created_at field.
	vehicleDescCreatedAt := vehicleMixinFields0[0].Descriptor()
	// vehicle.DefaultCreatedAt holds the default value on creation for the created_at field.
	vehicle.DefaultCreatedAt = vehicleDescCreatedAt.Default.(func() time.Time)
	// vehicleDescUpdatedAt is the schema descriptor for updated_at field.
	vehicleDescUpdatedAt := vehicleMixinFields0[1].Descriptor()
	// vehicle.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	vehicle.DefaultUpdatedAt = vehicleDescUpdatedAt.Default.(func() time.Time)
	// vehicle.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	vehicle.UpdateDefaultUpdatedAt = vehicleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// vehicleDescRegistrationNumber is the schema descriptor for registration_number field.
	vehicleDescRegistrationNumber := vehicleFields[0].Descriptor()
	// vehicle.RegistrationNumberValidator is a validator for the "registration_number" field. It is called by the builders before save.
	vehicle.RegistrationNumberValidator = vehicleDescRegistrationNumber.Validators[0].(func(string) error)
	// vehicleDescModel is the schema descriptor for model field.
	vehicleDescModel := vehicleFields[1].Descriptor()
	// vehicle.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	vehicle.ModelValidator = vehicleDescModel.Validators[0].(func(string) error)
	// vehicleDescSeat is the schema descriptor for seat field.
	vehicleDescSeat := vehicleFields[2].Descriptor()
	// vehicle.DefaultSeat holds the default value on creation for the seat field.
	vehicle.DefaultSeat = vehicleDescSeat.Default.(int)
	vehicleimageMixin := schema.VehicleImage{}.Mixin()
	vehicleimageMixinFields0 := vehicleimageMixin[0].Fields()
	_ = vehicleimageMixinFields0
	vehicleimageFields := schema.VehicleImage{}.Fields()
	_ = vehicleimageFields
	// vehicleimageDescCreatedAt is the schema descriptor for created_at field.
	vehicleimageDescCreatedAt := vehicleimageMixinFields0[0].Descriptor()
	// vehicleimage.DefaultCreatedAt holds the default value on creation for the created_at field.
	vehicleimage.DefaultCreatedAt = vehicleimageDescCreatedAt.Default.(func() time.Time)
	// vehicleimageDescUpdatedAt is the schema descriptor for updated_at field.
	vehicleimageDescUpdatedAt := vehicleimageMixinFields0[1].Descriptor()
	// vehicleimage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	vehicleimage.DefaultUpdatedAt = vehicleimageDescUpdatedAt.Default.(func() time.Time)
	// vehicleimage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	vehicleimage.UpdateDefaultUpdatedAt = vehicleimageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// vehicleimageDescImage is the schema descriptor for image field.
	vehicleimageDescImage := vehicleimageFields[0].Descriptor()
	// vehicleimage.ImageValidator is a validator for the "image" field. It is called by the builders before save.
	vehicleimage.ImageValidator = vehicleimageDescImage.Validators[0].(func(string) error)
}