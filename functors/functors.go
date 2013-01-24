package functors

//type FuncAppliesToString func(string) interface{}

func MapFuncOverStringStringHash(fx func(string) interface{}, mx map[string]string) map[string]interface{} {
    out := map[string]interface{}{}

    for i, val := range mx {
        out[i] = fx(val)
    }

    return out
}