package flygo

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//Define field struct
type Field struct {
	name string //parameter name

	/**preset part*/
	preset     bool   //set mode
	defaultVal string //default val parameter
	split      bool   //parameter split?
	splitRune  rune   //parameter split rune
	concat     bool   //parameter concat?
	concatRune rune   //parameter concat rune

	/**validate part*/
	validate  bool     //validate mode
	min       int      //min for number value
	max       int      //max for number value
	length    int      //length for string
	fixed     string   //fixed value for string
	minLength int      //min length for string
	maxLength int      //max length for string
	enums     []string //enums value for parameter
	regex     string   //regex test
}

//Define validateErr struct
type validateErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//new validateErr
func newValidateErr(msg string) *validateErr {
	return &validateErr{
		Code: app.GetValidateErrCode(),
		Msg:  msg,
	}
}

//json
func (validateErr *validateErr) json() string {
	bytes, err := json.Marshal(validateErr)
	if err != nil {
		app.log.fatal("json err : %v", err)
		return ""
	}
	return string(bytes)
}

func NewField(name string) *Field {
	return &Field{
		name:       name,
		preset:     true,
		defaultVal: "",
		split:      false,
		splitRune:  ',',
		concat:     false,
		concatRune: ',',
		validate:   true,
		length:     0,
		min:        0,
		max:        0,
		minLength:  0,
		maxLength:  0,
		fixed:      "",
		enums:      nil,
		regex:      "",
	}
}

//preset and validate
func presetAndValidate(fields []*Field, c *Context) error {
	if fields == nil || len(fields) <= 0 {
		return nil
	}
	var err error
	for _, field := range fields {
		presetField(field, c)
		err = validateField(field, c)
		if err != nil {
			break
		}
	}
	return err
}

//preset field
func presetField(field *Field, c *Context) {
	if !field.preset || field.name == "" {
		return
	}
	vals := c.Parameters[field.name]
	if vals == nil || len(vals) == 0 {
		//set default val
		vals = []string{field.defaultVal}
	}

	if field.split {
		param := vals[0]
		if !strings.ContainsRune(param, field.splitRune) {
			return
		}
		vals = strings.Split(param, fmt.Sprintf("%c", field.splitRune))
	}

	if field.concat {
		vals = []string{strings.Join(vals, fmt.Sprintf("%c", field.concatRune))}
	}
	c.Parameters[field.name] = vals
}

//validate field
func validateField(field *Field, c *Context) error {
	if !field.validate || field.name == "" {
		return nil
	}

	vals := c.Parameters[field.name]
	if vals == nil {
		return errors.New(newValidateErr(fmt.Sprintf("field[%v] is invalid", field.name)).json())
	}

	if field.min > 0 {
		for _, v := range vals {
			val, err := strconv.Atoi(v)
			if err != nil || val < field.min {
				return errors.New(newValidateErr(fmt.Sprintf("field[%v] is invalid with %v", field.name, "field.min")).json())
			}
		}
	}

	if field.max > 0 {
		for _, v := range vals {
			val, err := strconv.Atoi(v)
			if err != nil || val > field.max {
				return errors.New(newValidateErr(fmt.Sprintf("field[%v] is invalid with %v", field.name, "field.max")).json())
			}
		}
	}

	if field.length > 0 {
		for _, v := range vals {
			if len(v) != field.length {
				return errors.New(newValidateErr(fmt.Sprintf("field[%v] is invalid with %v", field.name, "field.length")).json())
			}
		}
	}

	if field.fixed != "" {
		for _, v := range vals {
			if v != field.fixed {
				return errors.New(newValidateErr(fmt.Sprintf("field[%v] is invalid with %v", field.name, "field.fixed")).json())
			}
		}
	}

	if field.minLength > 0 {
		for _, v := range vals {
			if len(v) < field.minLength {
				return errors.New(newValidateErr(fmt.Sprintf("field[%v] is invalid with %v", field.name, "field.minLength")).json())
			}
		}
	}

	if field.maxLength > 0 {
		for _, v := range vals {
			if len(v) > field.maxLength {
				return errors.New(newValidateErr(fmt.Sprintf("field[%v] is invalid with %v", field.name, "field.maxLength")).json())
			}
		}
	}

	if field.enums != nil && len(field.enums) > 0 {
		es := fmt.Sprintf("|%s|", strings.Join(field.enums, "|"))
		for _, v := range vals {
			if !strings.Contains(es, fmt.Sprintf("|%s|", v)) {
				return errors.New(newValidateErr(fmt.Sprintf("field[%v] is invalid with %v", field.name, "field.enums")).json())
			}
		}
	}

	if field.regex != "" {
		re := regexp.MustCompile(field.regex)
		for _, v := range vals {
			if !re.MatchString(v) {
				return errors.New(newValidateErr(fmt.Sprintf("field[%v] is invalid with %v", field.name, "field.regex")).json())
			}
		}
	}
	return nil
}

//set
func (field *Field) Preset(preset bool) *Field {
	field.preset = preset
	return field
}

//validate
func (field *Field) Validate(validate bool) *Field {
	field.validate = validate
	return field
}

//name
func (field *Field) Name(name string) *Field {
	field.name = name
	return field
}

// default val
func (field *Field) DefaultVal(defaultVal string) *Field {
	field.defaultVal = defaultVal
	return field
}

// min
func (field *Field) Min(min int) *Field {
	field.min = min
	return field
}

// max
func (field *Field) Max(max int) *Field {
	field.max = max
	return field
}

// length
func (field *Field) Length(length int) *Field {
	field.length = length
	return field
}

// fixed
func (field *Field) Fixed(fixed string) *Field {
	field.fixed = fixed
	return field
}

// minLength
func (field *Field) MinLength(minLength int) *Field {
	field.minLength = minLength
	return field
}

// maxLength
func (field *Field) MaxLength(maxLength int) *Field {
	field.maxLength = maxLength
	return field
}

// split
func (field *Field) Split(split bool) *Field {
	field.split = split
	return field
}

// split rune
func (field *Field) SplitRune(splitRune rune) *Field {
	field.splitRune = splitRune
	return field
}

// concat
func (field *Field) Concat(concat bool) *Field {
	field.concat = concat
	return field
}

// concat rune
func (field *Field) ConcatRune(concatRune rune) *Field {
	field.concatRune = concatRune
	return field
}

// enums
func (field *Field) Enums(enums ...string) *Field {
	field.enums = enums
	return field
}

// regex
func (field *Field) Regex(regex string) *Field {
	field.regex = regex
	return field
}
