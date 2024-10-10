import React from "react";
import { Link } from "react-router-dom";

export const About = () => {
  return (
    <>
      <h3>About</h3>
      <p>TaskList is an app that helps to keep track of pending tasks</p>
      <Link className="nav-link" to={"/"}>
        Return
      </Link>
    </>
  );
};