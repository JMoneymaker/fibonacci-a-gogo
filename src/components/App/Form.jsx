import React, { useState } from 'react';
import styles from './Form.css';
import { getFibonacci } from '../../services/getGoing';

const Form = () => {
  const [input, setInput] = useState('');
  const [output, setOutput] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();
    getFibonacci(input)
      .then(res => setOutput(res));
  };

  const handleChange = ({ target }) => {
    setInput(target.value);
  };

  return (
    <>
      <div className={styles.Fibonacci}>
        <form className={styles.Form} onSubmit={handleSubmit}>
          <input type='number' value={input} onChange={handleChange}></input>
          <button>Submit</button>
        </form>
        <div>{output}</div>
      </div>
    </>
  );
};

export default Form;
