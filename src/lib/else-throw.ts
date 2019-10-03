export const elseThrow = <T>(errorCb: () => Error): ((doc: T | null) => T) => (
  doc: T | null
): T => {
  if (!doc) {
    throw errorCb();
  }
  return doc;
};
