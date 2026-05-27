#version 140

in highp vec4 var_position;
in mediump vec3 var_normal;
in mediump vec2 var_texcoord0;
in mediump vec4 var_light;

out vec4 out_fragColor;

uniform mediump sampler2D tex0;

uniform fs_uniforms
{
	mediump vec4 tint;
};

void main()
{
	// Base color
	vec4 tint_pm = vec4(tint.rgb * tint.a, tint.a);
	vec4 color = texture(tex0, var_texcoord0.xy) * tint_pm;

	// Light vector
	vec3 L = normalize(var_light.xyz - var_position.xyz);
	float ndotl = max(dot(normalize(var_normal), L), 0.0);

	vec3 final_color;

	// Simple 2-band quantization: shadow vs main color
	if (ndotl < 0.1) {
		// Shadow band: slightly darker, desaturated
		float gray = dot(color.rgb, vec3(0.3,0.59,0.11));
		final_color = mix(vec3(gray), color.rgb, 0.2); // mostly gray
		final_color *= 0.6; // dim shadow
	} else {
		// Main flat color band
		final_color = color.rgb * 1.0; // no extra boost
	}

	out_fragColor = vec4(final_color, color.a);
}
