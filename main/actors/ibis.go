components {
  id: "player"
  component: "/main/actors/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Walk\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/actors/ibis/ibis.atlas\"\n"
  "}\n"
  ""
  position {
    y: 75.0
    z: 1.0
  }
  scale {
    x: 0.75
    y: 0.75
    z: 0.75
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -49.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"box\"\n"
  "  }\n"
  "  data: 72.5\n"
  "  data: 48.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
