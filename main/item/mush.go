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
  "  texture: \"/assets/items/mush/mush.atlas\"\n"
  "}\n"
  ""
  position {
    y: -17.0
  }
  scale {
    x: 0.5
    y: 0.5
    z: 0.5
  }
}
