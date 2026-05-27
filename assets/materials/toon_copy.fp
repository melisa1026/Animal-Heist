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
	vec4 tint_pm = vec4(tint.rgb * tint.a, tint.a);
	vec4 color = texture(tex0, var_texcoord0.xy) * tint_pm;

	vec3 L = normalize(var_light.xyz - var_position.xyz);
	float ndotl = max(dot(normalize(var_normal), L), 0.0);

	vec3 final_color;

	if (ndotl < 0.15) {
		// Shadows: desaturated & dark
		float gray = dot(color.rgb, vec3(0.3,0.59,0.11));
		final_color = mix(vec3(gray), color.rgb, 0.3);
		final_color *= 0.5;
	} else if (ndotl < 0.55) {
		// Midtones: pop in color but slightly darker than highlight
		vec3 saturated = pow(color.rgb, vec3(0.7)); // boost color
		final_color = saturated * 1.2;             // moderate brightness
	} else {
		// Highlights: brightest, normal saturation
		final_color = color.rgb * 2.0;             // strong brightness, no extra saturation
	}

	out_fragColor = vec4(final_color, color.a);
}
