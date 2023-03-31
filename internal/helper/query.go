package helper

import (
	"fmt"
	"reflect"
	"strconv"
)

func QueryWhere(w interface{}) (string, []interface{}) {
	var (
		wheres = []string{}
		vals   []interface{}
	)

	query := ``

	t := reflect.TypeOf(w)
	v := reflect.ValueOf(w)

	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).IsZero() != true {
			wheres = append(wheres, t.Field(i).Tag.Get("db"))
			vals = append(vals, v.Field(i).Interface())
		}
	}

	for i, p := range wheres {
		if i > 0 {
			query += fmt.Sprintf(`AND %s = $%v `, p, i+1)
		} else {
			query += fmt.Sprintf(`WHERE %s = $%v `, p, i+1)

		}
	}

	return query, vals
}

func QueryInsert(w interface{}) (string, []interface{}) {
	var (
		fields = []string{}
		vals   []interface{}
	)

	queryFields := ``
	queryParams := ``

	t := reflect.TypeOf(w)
	v := reflect.ValueOf(w)

	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).IsZero() != true {
			fields = append(fields, t.Field(i).Tag.Get("db"))
			vals = append(vals, v.Field(i).Interface())
		}
	}

	for i, v := range fields {
		if i > 0 {
			queryFields += fmt.Sprintf(`, %s`, v)
			queryParams += fmt.Sprintf(`, $%s`, strconv.Itoa(i+1))
		} else {
			queryFields += fmt.Sprintf(`%s`, v)
			queryParams += fmt.Sprintf(`$%s`, strconv.Itoa(i+1))

		}
	}

	result := fmt.Sprintf(`(%s) VALUES (%s)`, queryFields, queryParams)

	return result, vals
}

func QueryUpdate(p, w interface{}) (string, []interface{}) {
	var (
		vals []interface{}
	)

	query := ``

	//UPDATE table_name
	//SET column1 = value1, column2 = value2, ...
	//WHERE condition;

	typeParam := reflect.TypeOf(p)
	valParam := reflect.ValueOf(p)

	for i := 0; i < typeParam.NumField(); i++ {
		if valParam.Field(i).IsZero() != true {
			if query != `` {
				query += fmt.Sprintf(`, %s = '%s'`, typeParam.Field(i).Tag.Get("db"), valParam.Field(i).Interface())
			} else {
				query += fmt.Sprintf(`%s = '%s'`, typeParam.Field(i).Tag.Get("db"), valParam.Field(i).Interface())
			}

		}
	}

	//Where
	typeWhere := reflect.TypeOf(w)
	valWhere := reflect.ValueOf(w)

	for i := 0; i < typeWhere.NumField(); i++ {
		if valWhere.Field(i).IsZero() != true {

			vals = append(vals, valWhere.Field(i).Interface())

			if i > 0 {
				query += fmt.Sprintf(`AND %s = $%v `, typeWhere.Field(i).Tag.Get("db"), i+1)
			} else {
				query += fmt.Sprintf(` WHERE %s = $%v `, typeWhere.Field(i).Tag.Get("db"), i+1)

			}
		}
	}

	return query, vals
}
