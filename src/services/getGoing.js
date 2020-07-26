export const getFibonacci = userInput => {
  return fetch(`http://localhost:8080/api/hello/${userInput}`)
    .then(res => res.text());
};
