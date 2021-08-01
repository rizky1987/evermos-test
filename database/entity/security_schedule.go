package entity

import (
	"encoding/json"
	"evermos-test/helper"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"time"
)

type SecuritySchedule struct {
	Id           				*bson.ObjectId 			`bson:"_id,omitempty" json:"_id"`
	RoundingSequence 			int               		`bson:"rounding_sequence" json:"rounding_sequence"`
	SecurityGroupId  			bson.ObjectId 			`bson:"security_group_id" json:"security_group_id"`
	SecurityGroups  			[]*SecurityGroup		`bson:"temp_security_groups,omitempty" json:"security_groups,omitempty"`
	ScheduleDate     			time.Time         		`bson:"schedule_date" json:"schedule_date"`
	UserAttenders    			[]*UserAttender   		`bson:"user_attenders"`
	CreatedAt        			time.Time         		`bson:"created_at"`
	CreatedAtUTC     			time.Time         		`bson:"created_at_utc"`
	UpdatedAt        			*time.Time        		`bson:"updated_at,omitempty"`
	UpdatedAtUTC     			*time.Time        		`bson:"updated_at_utc,omitempty"`
}

type UserAttender struct {
	UserId           			bson.ObjectId 			`bson:"user_id"`
	IsAttended           		bool					`bson:"is_attended"`
	IsExceptionalAttend    		bool					`bson:"is_exceptional_attend"`
	AttendingTime    		    string					`bson:"attending_time"`
	Explanation           		string					`bson:"explanation,omitempty"`
	Image           			string					`bson:"image,omitempty"`
}

// Begin Create Validation

func (entityStruct *SecuritySchedule) GenerateTransactionData() error {

	entityStruct.CreatedAt = helper.GetCurrentTimeAsiaJakarta()
	entityStruct.CreatedAtUTC = helper.GetCurrentTimeUTC()

	return nil
}

//Fungsi ini masih dalam percobaan jd jika masih ada yg blm termappingkan silahkan mapping manual
func (entityStruct *SecuritySchedule) MappingCreateDataToEntityStruct(requestedStruct request.CreateSecurityScheduleRequest) {

	jsonString, _ := json.Marshal(requestedStruct)
	json.Unmarshal(jsonString, &entityStruct)

	//remove all white space in right and left string field type, make sure we doesn't save unused data
	entityStruct.TrimStringFieldType()

}

// End Create validation

// Begin Mutual Function For Create and Update
func (entityStruct *SecuritySchedule) TrimStringFieldType() {

	msValuePtr := reflect.ValueOf(entityStruct)
	msValue := msValuePtr.Elem()

	for i := 0; i < msValue.NumField(); i++ {
		field := msValue.Field(i)

		// Ignore fields that don't have the same type as a string
		if field.Type() != reflect.TypeOf("") {
			continue
		}

		str := field.Interface().(string)
		str = helper.TrimWhiteSpace(str)
		field.SetString(str)
	}
}

// Please make sure all data is a correct data before we save it to DB
func (entityStruct *SecuritySchedule) ValidatebeforeSaveToDB() error {

	entityStruct.TrimStringFieldType()
	// Add you Validation Here

	return nil
}

// End Mutual Function For Create and Update