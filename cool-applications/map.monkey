let map = fn(arr, f){ let iter = fn(arr, accumulated){ if (len(arr) == 0) {return accumulated} else {return iter(rest(arr), push(accumulated, f(first(arr))))} } return iter(arr, []); }
