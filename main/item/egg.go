components {
  id: "wiggle"
  component: "/main/item/wiggle.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Wiggle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/items/egg/egg.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
}
