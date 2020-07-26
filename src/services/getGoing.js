export const getFibonacci = userInput => {
  return fetch(`http://localhost:8080/api/fibonacci/${userInput}`)
    .then(res => res.json());
};
