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
      <form className={styles.Form} onSubmit={handleSubmit}>
        <input type='text' value={input} onChange={handleChange}></input>
        <button>Submit</button>
      </form>
      <div>{output}</div>
    </>
  );
};

export default Form;
