export function getFileBaseName(name: string) {
  return name.split(/\.(?=[^.]+$)/)[0]
}

export default getFileBaseName
