components {
  id: "BubbleClicker"
  component: "/Scripts/BubbleClicker.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "minHeight"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "maxHeight"
    value: "500.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "minWidth"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "maxWidth"
    value: "800.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "power_up_factory"
  component: "/main/power_up_factory.factory"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/Assets/game.atlas\"\ndefault_animation: \"bubble\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "clickSound"
  type: "sound"
  data: "sound: \"/Assets/soundFX/clickSound.wav\"\nlooping: 0\ngroup: \"master\"\ngain: 0.3\npan: 0.0\nspeed: 1.0\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
