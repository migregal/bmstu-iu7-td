export function errorsObjectToArray(errors: Record<string, string>, knownFields: string[] = [], onlyKnown = false) {
  errors = {...errors}
  const arr = []

  for (const key of knownFields) {
    if (key in errors) {
      arr.push(errors[key])
      delete errors[key]
    }
  }

  return onlyKnown ? arr : arr.concat(Object.values(errors))
}

export default errorsObjectToArray
