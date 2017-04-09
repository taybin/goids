'use strict';

import BABYLON from 'babylonjs';
import boidProto from '../boid.proto';

const BoidMessage = boidProto.main.BoidPosition;

window.addEventListener("DOMContentLoaded", evt => {
  const boids = {}

  const canvas = document.getElementById('renderCanvas');
  const engine = new BABYLON.Engine(canvas, true);

  const createScene = () => {
    // create a basic BJS Scene object
    const scene = new BABYLON.Scene(engine);

    // create a FreeCamera, and set its position to (x:0, y:5, z:-10)
    const camera = new BABYLON.FreeCamera('camera1', new BABYLON.Vector3(0, 5, -500), scene);

    // target the camera to scene origin
    camera.setTarget(BABYLON.Vector3.Zero());

    // attach the camera to the canvas
    camera.attachControl(canvas, false);

    // create a basic light, aiming 0,1,0 - meaning, to the sky
    const light = new BABYLON.HemisphericLight('light1', new BABYLON.Vector3(0,5,-500), scene);

    // return the created scene
    return scene;
  };

  const scene = createScene();

  engine.runRenderLoop( () => {
    scene.render();
  });

  window.addEventListener('resize', () => {
    engine.resize();
  });

  const ws = new WebSocket("ws://localhost:3000/ws");
  ws.binaryType = 'arraybuffer';
  ws.onopen = evt => {
    console.log("OPEN");
  }

  ws.onclose = evt => {
    console.log("CLOSE");
    ws = null;
  }

  ws.onerror = evt => {
    console.log("ERROR: " + evt.data);
  }

  ws.onmessage = evt => {
    const boid = BoidMessage.decode(new Uint8Array(evt.data));

    if (!(boid.id in boids)) {
      const cone = BABYLON.MeshBuilder.CreateCylinder(boid.id.toString(), {
        diameterTop: 0,
        diameterBottom: 3,
        tessellation: 4,
        height: 4
      }, scene);
      cone.position =
        new BABYLON.Vector3(
          boid.position[0],
          boid.position[1],
          boid.position[2]);
      boid.cone = cone;
      boids[boid.id] = boid;
    } else {
      const newPosition = new BABYLON.Vector3(
        boid.position[0],
        boid.position[1],
        boid.position[2]
      );
      BABYLON.Animation.CreateAndStartAnimation(
        'animation-'+boid.id+'-position',
        boids[boid.id].cone,
        'position',
        30,
        15,
        boids[boid.id].cone.position,
        newPosition,
        0,
        null
      );
      boids[boid.id].cone.lookAt(newPosition,
        0,
        (Math.PI / 2) * 3
      );
    }
  };
});
