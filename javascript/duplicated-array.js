// given the array [1, 2, 5, 3, 1, 45, 13, 3, 22, 4, 5]
// return the array in an orderly manner and without duplicates
// example: [1, 2, 3, 4, 5, 13, 22, 45]

const array = [1, 2, 5, 3, 1, 45, 13, 3, 22, 4, 5]

const setWithoutDuplicity = new Set(array)
const ordelyArray =  Array.from(setWithoutDuplicity).sort((a, b) => a - b)

console.log(ordelyArray) // [1, 2, 3, 4, 5, 13, 22, 45]
