const FILE_SIZES = [
  "b",
  "kb",
  "mb",
  "gb",
]

export function pluralizeFileSize(size: number) {
  let i = 0
  while (i < FILE_SIZES.length - 1 && Math.ceil(size) > 1024) {
    size /= 1024
    i++
  }
  return Math.ceil(size) + FILE_SIZES[i]
}

export default pluralizeFileSize
