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
	name        string //parameter name
	description string //parameter description

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
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	logger *log
}

//new validateErr
func newValidateErr(code int, msg string, logger *log) *validateErr {
	return &validateErr{
		Code:   code,
		Msg:    msg,
		logger: logger,
	}
}

//json
func (validateErr *validateErr) json() string {
	bytes, err := json.Marshal(validateErr)
	if err != nil {
		validateErr.logger.Error("json err : %v", err)
		return ""
	}
	return string(bytes)
}

func NewField(name string) *Field {
	return &Field{
		name:        name,
		description: name,
		preset:      true,
		defaultVal:  "",
		split:       false,
		splitRune:   ',',
		concat:      false,
		concatRune:  ',',
		validate:    true,
		length:      0,
		min:         0,
		max:         0,
		minLength:   0,
		maxLength:   0,
		fixed:       "",
		enums:       nil,
		regex:       "",
	}
}

//preset and validate
func (c *Context) presetAndValidate(fields []*Field) error {
	if fields == nil || len(fields) <= 0 {
		return nil
	}
	var err error
	for _, field := range fields {
		c.presetField(field)
		err = c.validateField(field)
		if err != nil {
			break
		}
	}
	return err
}

//preset field
func (c *Context) presetField(field *Field) {
	if !field.preset || field.name == "" {
		return
	}
	vals := c.ParamMap[field.name]
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
	c.ParamMap[field.name] = vals
}

//validate field
func (c *Context) validateField(field *Field) error {
	code := c.app.Config.Flygo.Validate.Code
	logger := c.app.Logger
	if !field.validate || field.name == "" {
		return nil
	}

	errFn := func(msg string) error {
		return errors.New(newValidateErr(code, msg, logger).json())
	}

	vals := c.ParamMap[field.name]
	if vals == nil {
		return errFn(fmt.Sprintf("field[%v] validated[null] fail", field.name))
	}

	if field.min > 0 {
		for _, v := range vals {
			val, err := strconv.Atoi(v)
			if err != nil || val < field.min {
				return errFn(fmt.Sprintf("field[%v] validated[%v] fail", field.name, "field.min"))
			}
		}
	}

	if field.max > 0 {
		for _, v := range vals {
			val, err := strconv.Atoi(v)
			if err != nil || val > field.max {
				return errFn(fmt.Sprintf("field[%v] validated[%v] fail", field.name, "field.max"))
			}
		}
	}

	if field.length > 0 {
		for _, v := range vals {
			if len(v) != field.length {
				return errFn(fmt.Sprintf("field[%v] validated[%v] fail", field.name, "field.length"))
			}
		}
	}

	if field.fixed != "" {
		for _, v := range vals {
			if v != field.fixed {
				return errFn(fmt.Sprintf("field[%v] validated[%v] fail", field.name, "field.fixed"))
			}
		}
	}

	if field.minLength > 0 {
		for _, v := range vals {
			if len(v) < field.minLength {
				return errFn(fmt.Sprintf("field[%v] validated[%v] fail", field.name, "field.minLength"))
			}
		}
	}

	if field.maxLength > 0 {
		for _, v := range vals {
			if len(v) > field.maxLength {
				return errFn(fmt.Sprintf("field[%v] validated[%v] fail", field.name, "field.maxLength"))
			}
		}
	}

	if field.enums != nil && len(field.enums) > 0 {
		es := fmt.Sprintf("|%s|", strings.Join(field.enums, "|"))
		for _, v := range vals {
			if !strings.Contains(es, fmt.Sprintf("|%s|", v)) {
				return errFn(fmt.Sprintf("field[%v] validated[%v] fail", field.name, "field.enums"))
			}
		}
	}

	if field.regex != "" {
		re := regexp.MustCompile(field.regex)
		for _, v := range vals {
			if !re.MatchString(v) {
				return errFn(fmt.Sprintf("field[%v] validated[%v] fail", field.name, "field.regex"))
			}
		}
	}
	return nil
}

//description
func (field *Field) Description(description string) *Field {
	field.description = description
	return field
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
