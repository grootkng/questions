// Findout if a word is a palindrome or not
// returning true or false
// example: Omissíssimo, Aibofobia

function main(...args) {
  const word = args[0].toLowerCase()
  const reverseWord = word.split('').reverse().join('')

  if (reverseWord == word) {
    return true
  } else {
    return false
  }
}

console.log(main('Omissíssimo')) // true
console.log(main('Aibofobia')) // true
console.log(main('kayak')) // true
console.log(main('madam')) // true
console.log(main('world')) // false
