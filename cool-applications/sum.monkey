let reduce = fn(arr, initial, f){ let iter = fn(arr, result){ if (len(arr) == 0) {return result} else {iter(rest(arr), f(result, first(arr)))} } return iter(arr, initial)};
let sum = fn(arr) { reduce(arr, 0, fn(initial, el){initial + el})};
