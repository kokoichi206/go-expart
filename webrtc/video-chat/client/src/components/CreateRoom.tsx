import React from "react";
import { RouteComponentProps } from 'react-router-dom';

const CreateRoom: React.FC<RouteComponentProps>  = (props) => {
  const create = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();

    const resp = await fetch("http://localhost:8080/create");
    const { room_id } = await resp.json();
    console.log(`got roomID: ${room_id}`)

    props.history.push(`/room/${room_id}`);
  };

  return (
    <div>
      <button onClick={create}>Create Room</button>
    </div>
  );
};

export default CreateRoom;
