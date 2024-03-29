// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package execs

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/4x2Hach1/learning/gql-go/api/models"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Memory_id(ctx context.Context, field graphql.CollectedField, obj *models.Memory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Memory_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Memory_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Memory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Memory_userId(ctx context.Context, field graphql.CollectedField, obj *models.Memory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Memory_userId(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UserID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Memory_userId(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Memory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Memory_title(ctx context.Context, field graphql.CollectedField, obj *models.Memory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Memory_title(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Title, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Memory_title(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Memory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Memory_date(ctx context.Context, field graphql.CollectedField, obj *models.Memory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Memory_date(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Date, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNDateTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Memory_date(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Memory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type DateTime does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Memory_location(ctx context.Context, field graphql.CollectedField, obj *models.Memory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Memory_location(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Location, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Memory_location(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Memory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Memory_detail(ctx context.Context, field graphql.CollectedField, obj *models.Memory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Memory_detail(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Detail, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Memory_detail(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Memory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputDeleteMemory(ctx context.Context, obj interface{}) (models.DeleteMemory, error) {
	var it models.DeleteMemory
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalNInt2int(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputNewMemory(ctx context.Context, obj interface{}) (models.NewMemory, error) {
	var it models.NewMemory
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"title", "date", "location", "detail"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "title":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("title"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Title = data
		case "date":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("date"))
			data, err := ec.unmarshalNDateTime2timeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.Date = data
		case "location":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("location"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Location = data
		case "detail":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("detail"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Detail = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputUpdateMemory(ctx context.Context, obj interface{}) (models.UpdateMemory, error) {
	var it models.UpdateMemory
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "title", "date", "location", "detail"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalNInt2int(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "title":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("title"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Title = data
		case "date":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("date"))
			data, err := ec.unmarshalNDateTime2timeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.Date = data
		case "location":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("location"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Location = data
		case "detail":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("detail"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Detail = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var memoryImplementors = []string{"Memory"}

func (ec *executionContext) _Memory(ctx context.Context, sel ast.SelectionSet, obj *models.Memory) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, memoryImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Memory")
		case "id":
			out.Values[i] = ec._Memory_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "userId":
			out.Values[i] = ec._Memory_userId(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "title":
			out.Values[i] = ec._Memory_title(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "date":
			out.Values[i] = ec._Memory_date(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "location":
			out.Values[i] = ec._Memory_location(ctx, field, obj)
		case "detail":
			out.Values[i] = ec._Memory_detail(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNDeleteMemory2githubᚗcomᚋ4x2Hach1ᚋlearningᚋgqlᚑgoᚋapiᚋmodelsᚐDeleteMemory(ctx context.Context, v interface{}) (models.DeleteMemory, error) {
	res, err := ec.unmarshalInputDeleteMemory(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNMemory2ᚕᚖgithubᚗcomᚋ4x2Hach1ᚋlearningᚋgqlᚑgoᚋapiᚋmodelsᚐMemoryᚄ(ctx context.Context, sel ast.SelectionSet, v []*models.Memory) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNMemory2ᚖgithubᚗcomᚋ4x2Hach1ᚋlearningᚋgqlᚑgoᚋapiᚋmodelsᚐMemory(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNMemory2ᚖgithubᚗcomᚋ4x2Hach1ᚋlearningᚋgqlᚑgoᚋapiᚋmodelsᚐMemory(ctx context.Context, sel ast.SelectionSet, v *models.Memory) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Memory(ctx, sel, v)
}

func (ec *executionContext) unmarshalNNewMemory2githubᚗcomᚋ4x2Hach1ᚋlearningᚋgqlᚑgoᚋapiᚋmodelsᚐNewMemory(ctx context.Context, v interface{}) (models.NewMemory, error) {
	res, err := ec.unmarshalInputNewMemory(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNUpdateMemory2githubᚗcomᚋ4x2Hach1ᚋlearningᚋgqlᚑgoᚋapiᚋmodelsᚐUpdateMemory(ctx context.Context, v interface{}) (models.UpdateMemory, error) {
	res, err := ec.unmarshalInputUpdateMemory(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOMemory2ᚖgithubᚗcomᚋ4x2Hach1ᚋlearningᚋgqlᚑgoᚋapiᚋmodelsᚐMemory(ctx context.Context, sel ast.SelectionSet, v *models.Memory) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Memory(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
