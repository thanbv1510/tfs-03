// Ex1

// Ex2
function isPalindrome(str) {
    let tempStr = ""
    // remove special characters
    str = str.toLowerCase()
    for (let char of str) {
        if (char >= 'a' && char <= 'z') {
            tempStr += char
        }
    }

    // check is palindrome
    for (let i = 0; i <= tempStr.length / 2; i++) {
        if (tempStr[i] !== tempStr[tempStr.length - 1 - i]) {
            return false
        }
    }

    return true
}


// Ex3
function uniqueUnion(...params) {
    let result = new Set()
    params.forEach(arr => result = new Set([...result, ...arr]))

    return Array.from(result)
}

// Ex4
function seekAndDestroy(dataArr, ...params) {
    let paramArr = Array.from(params)

    return dataArr.filter(item => !paramArr.includes(item))
}

// Ex5

// Ex6
function drop(dataArr, somethingFunc) {
    let index = 1
    while (!somethingFunc(index)) {
        dataArr.shift()
        index++
    }

    return dataArr
}
