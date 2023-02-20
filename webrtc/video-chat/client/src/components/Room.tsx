import React, { useEffect, useRef } from "react";
import { RouteComponentProps } from "react-router-dom";

const Room = (props) => {
  const userVideo = useRef();
  const userStream = useRef();
  const partnerVideo = useRef();
  const peerRef = useRef();
  const webSocketRef = useRef();

  const openCamera = async () => {
    const allDevices = await navigator.mediaDevices.enumerateDevices();
    console.log("allDevices: ");
    console.log(allDevices);
    const cameras = allDevices.filter((device) => device.kind == "videoinput");
    console.log("cameras: ");
    console.log(cameras);

    const constraints = {
      audio: true,
      video: {
        // 1 かどうかは laptop かどうか関係？
        deviceId: cameras[0].deviceId,
      },
    };

    try {
      return await navigator.mediaDevices.getUserMedia(constraints);
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    openCamera().then((stream) => {
      userVideo.current.srcObject = stream;
      userStream.current = stream;

      const id = props.match.params.roomID;
      if (!webSocketRef.current) {
        webSocketRef.current = new WebSocket(
          `ws://localhost:8080/join?roomID=${id}`
        );
      }

      webSocketRef.current.addEventListener("open", () => {
        webSocketRef.current.send(JSON.stringify({ join: true }));
      });

      webSocketRef.current.addEventListener("message", async (e) => {
        console.log("e: ");
        console.log(e);
        const message = JSON.parse(e.data);

        if (message.join) {
          // start RTC process
          callUser();
        }

        if (message.offer) {
          handleOffer(message.offer);
        }

        if (message.answer) {
          console.log("Receiving answer");
          peerRef.current.setRemoteDescription(
            new RTCSessionDescription(message.answer)
          );
        }

        if (message.iceCandidate) {
          console.log("Receiving and adding ice candidate");
          try {
            await peerRef.current.addIceCandidate(message.iceCandidate);
          } catch (e) {
            console.log("Error when receiving ice candidate: ", e);
          }
        }
      });
    });
  });

  const handleOffer = async (offer) => {
    console.log("Received offer, creating answer");

    peerRef.current = createPeer();

    await peerRef.current.setRemoteDescription(
      new RTCSessionDescription(offer)
    );

    userStream.current.getTracks().forEach((track) => {
      peerRef.current.addTrack(track, userStream.current);
    });

    const answer = await peerRef.current.createAnswer();
    await peerRef.current.setLocalDescription(answer);

    webSocketRef.current.send(
      JSON.stringify({ answer: peerRef.current.localDescription })
    );
  };

  const callUser = () => {
    console.log("Calling other user");
    peerRef.current = createPeer();

    userStream.current.getTracks().forEach((track) => {
      peerRef.current.addTrack(track, userStream.current);
    });
  };

  const createPeer = () => {
    console.log("Creating peer connection");

    const peer = new RTCPeerConnection({
      iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
    });

    peer.onnegotiationneeded = handleNegotiationNeeded;
    peer.onicecandidate = handleIceCandidateEvent;
    peer.ontrack = handleTrackEvent;

    return peer;
  };

  const handleNegotiationNeeded = () => {
    console.log("Creating offer...");

    try {
      const myOffer = peerRef.current.createOffer();
      peerRef.current.setLocalDescription(myOffer);

      webSocketRef.current.send(
        JSON.stringify({ offer: peerRef.current.localDescription })
      );
    } catch (e) {
      console.log("Error when ");
    }
  };
  const handleIceCandidateEvent = (e) => {
    console.log("Found ice candidate");
    if (e.candidate) {
      console.log(e.candidate);
      webSocketRef.current.send(JSON.stringify({ iceCandidate: e.candidate }));
    }
  };
  const handleTrackEvent = (e) => {
    console.log("Received tracks");
    partnerVideo.current.srcObject = e.streams[0];
  };

  return (
    <div>
      <video autoPlay controls={true} ref={userVideo}></video>
      <video autoPlay controls={true} ref={partnerVideo}></video>
    </div>
  );
};

export default Room;
