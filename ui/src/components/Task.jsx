import React, { useState } from "react";
import { TrashCanIcon} from "./TrashCanSvg";

export const Task = (props) => {
  const [completed, setCompleted] = useState(props.task.completed);
  const [showDelete, setShowDelete] = useState(false);

  async function handleCheckClick(e) {
    const payload = {
      completed: e.target.checked,
    }

    const resp = await fetch(`http://localhost:8080/tasks/${props.task.id}`, {
      method: "PUT",
      body: JSON.stringify(payload),
    });

    const body = await resp.json();
    const { task } = body;
    setCompleted(task.completed);
  }

  async function handleDelete(e) {
    await fetch(`http://localhost:8080/tasks/${props.task.id}`, {
      method: "DELETE" });
    props.onDeleteSuccess();
  }

  return (
    <div
      className="task"
      onMouseEnter={()=> setShowDelete(true)}
      onMouseLeave={()=> setShowDelete(false)}
      >
      <input
        type={"checkbox"}
        className="checkbox"
        checked={completed}
        onChange={handleCheckClick}
      />
      <p>{props.task.description}</p>
      {showDelete && (
        <button className="delete" onClick={handleDelete}>
          <TrashCanIcon />
        </button>
      )}
    </div>
  );
};