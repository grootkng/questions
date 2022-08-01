// --- Anagram
// Check to see if two provided strings are anagrams of each other.
// One string is an anagram of another if it uses the same characters
// in the same quantity. Only consider characters, not spaces
// or punctuation.  Consider capital letters to be the same as lower case

// --- Examples
//   anagrams('rail safety', 'fairy tales') --> True
//   anagrams('RAIL! SAFETY!', 'fairy tales') --> True
//   anagrams('Hi there', 'Bye there') --> False

const specialChars = /[^\w\s]/g // match everything that is not a word or a whitespace

function anagrams(...args) {
  const paramA = args[0]
  const paramB = args[1]

  const a = paramA.replace(specialChars, '').toLowerCase()
  const b = paramB.replace(specialChars, '').toLowerCase()

  if(a.length !== b.length){
     return false
  }

  const strA = a.split('').sort().join('')
  const strB = b.split('').sort().join('')

  return strA === strB ? true : false
}

console.log(anagrams('rail safety', 'fairy tales'))
console.log(anagrams('RAIL! SAFETY!', 'fairy tales'))
console.log(anagrams('Hi there', 'Bye there'))
