import BABYLON from 'babylonjs';
import boidProto from '../boid.proto';

var BoidMessage = boidProto.main.BoidPosition;

window.addEventListener("DOMContentLoaded", function(evt) {
  var boids = {}

  var canvas = document.getElementById('renderCanvas');
  var engine = new BABYLON.Engine(canvas, true);

  var createScene = function() {
    // create a basic BJS Scene object
    var scene = new BABYLON.Scene(engine);

    // create a FreeCamera, and set its position to (x:0, y:5, z:-10)
    var camera = new BABYLON.FreeCamera('camera1', new BABYLON.Vector3(0, 5, -500), scene);

    // target the camera to scene origin
    camera.setTarget(BABYLON.Vector3.Zero());

    // attach the camera to the canvas
    camera.attachControl(canvas, false);

    // create a basic light, aiming 0,1,0 - meaning, to the sky
    var light = new BABYLON.HemisphericLight('light1', new BABYLON.Vector3(0,1,0), scene);

    // return the created scene
    return scene;
  };

  var scene = createScene();

  engine.runRenderLoop(function() {
    scene.render();
  });

  window.addEventListener('resize', function() {
    engine.resize();
  });

  var ws;
  ws = new WebSocket("ws://localhost:3000/ws");
  ws.binaryType = 'arraybuffer';
  ws.onopen = function(evt) {
    console.log("OPEN");
  }

  ws.onclose = function(evt) {
    console.log("CLOSE");
    ws = null;
  }

  ws.onerror = function(evt) {
    console.log("ERROR: " + evt.data);
  }

  ws.onmessage = function(evt) {
    var boid = BoidMessage.decode(new Uint8Array(evt.data));

    if (!(boid.id in boids)) {
      var cone = BABYLON.MeshBuilder.CreateCylinder(boid.id.toString(), {
        diameterTop: 0,
        diameterBottom: 3,
        tessellation: 4,
        height: 4
      }, scene);
      boid.cone = cone;
      boids[boid.id] = boid;
      boids[boid.id].cone.position =
        new BABYLON.Vector3(
          boid.position[0],
          boid.position[1],
          boid.position[2]);
    } else {
      BABYLON.Animation.CreateAndStartAnimation(
        'animation-'+boid.id+'-position',
        boids[boid.id].cone,
        'position',
        30,
        120,
        boids[boid.id].cone.position,
        new BABYLON.Vector3(
          boid.position[0],
          boid.position[1],
          boid.position[2]),
        0,
        null
      );
    }
  };
});
