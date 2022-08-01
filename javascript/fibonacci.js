// Sequence is a series of numbers in which the next number is gotten by adding up the two numbers before it
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
// Write a method that when passed a number returns the Fibonacci sequence with a length of the number

const MAX_NUMBER = 10

function main(number) {
  return new Array(number)
    .fill(1) // fill the array with length MAX_NUMBER with 1
    .reduce((arr, _ , i) => {
      if (i <= 1) { // for 0 or 1
        arr.push(i)
      } else {
        arr.push(arr[i-2] + arr[i-1])
      }
      return arr
    },[])
}

console.log(main(MAX_NUMBER))
