components {
  id: "cam-click"
  component: "/main/camera/cam-click.script"
}
components {
  id: "rts-camera"
  component: "/main/camera/rts-camera.script"
}
embedded_components {
  id: "camera"
  type: "camera"
  data: "aspect_ratio: 1.0\n"
  "fov: 0.7854\n"
  "near_z: 1.0\n"
  "far_z: 5000.0\n"
  "auto_aspect_ratio: 1\n"
  "orthographic_projection: 1\n"
  "orthographic_zoom: 96.0\n"
  ""
}
embedded_components {
  id: "tag"
  type: "factory"
  data: "prototype: \"/main/camera/tag.go\"\n"
  ""
}
