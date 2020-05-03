function staircase(n) {
  for (let i = n -1; i >= 0; i--) {
    const hashesToPrint = n - i
    const spacesToPrint = n - hashesToPrint

    const hashes = Array(hashesToPrint).fill("#").join("")
    const spaces = Array(spacesToPrint).fill(" ").join("")

    console.log(`${spaces}${hashes}`)
  }
}


staircase(6)
