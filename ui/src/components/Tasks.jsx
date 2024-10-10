import React, {useCallback, useEffect, useState} from "react";
import {Task} from "./Task";
import {Link} from "react-router-dom";
import {NewTaskInput} from "./NewTaskForm";

export const Tasks = () => {

  const [tasks,setTasks] = useState([]);

  const fetchTasks = useCallback(async () => {
    const resp = await fetch ("http://localhost:8080/tasks");
     const body = await resp.json();
     // console.log(body);

    const { tasks } = body;
    setTasks(tasks);
  }, [setTasks]);

  useEffect(() => {
    fetchTasks();
  }, [fetchTasks]);

  function onDeleteSuccess() {
    fetchTasks();
  }

  function onCreateSuccess(newTask) {
    setTasks([...tasks, newTask])
  }

  return (
    <>
      <h3> To Do:</h3>
      <div className="tasks">
        {tasks.map((task) => (
          <Task key={task.id} task={task} onDeleteSuccess={onDeleteSuccess} />
        ))}
      </div>
      <NewTaskInput onCreateSuccess={onCreateSuccess} />
      <Link to="/about" className="nav-link">
        Learn More....
      </Link>
    </>
  );
};

