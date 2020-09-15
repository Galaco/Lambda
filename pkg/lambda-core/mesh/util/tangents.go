package util

import "github.com/go-gl/mathgl/mgl32"

// GenerateTangents generates tangents for vertex data
func GenerateTangents(points []float32, normals []float32, texCoords []float32) (tangents []float32) {
	//const vector<vec3> & points,
	//const vector<vec3> & normals,
	//const vector<int> & faces,
	//const vector<vec2> & texCoords,
	//	vector<vec4> & tangents)
	//{
	//vector<vec3> tan1Accum;
	tan1Accum := make([]float32, len(points))
	//vector<vec3> tan2Accum;
	tan2Accum := make([]float32, len(points))
	tangents = make([]float32, len(points)+(len(points)/3))

	//for( uint i = 0; i < points.size(); i++ ) {
	//tan1Accum.push_back(vec3(0.0f));
	//tan2Accum.push_back(vec3(0.0f));
	//tangents.push_back(vec4(0.0f));
	//}

	// Compute the tangent vector
	for i := uint(0); i < uint(len(points))-9; i += 9 {
		rootIdx := i / 3
		p1 := mgl32.Vec3{points[i], points[i+1], points[i+2]}
		p2 := mgl32.Vec3{points[i+3], points[i+4], points[i+5]}
		p3 := mgl32.Vec3{points[i+6], points[i+7], points[i+8]}
		//const vec3 &p1 = points[faces[i]];
		//const vec3 &p2 = points[faces[i+1]];
		//const vec3 &p3 = points[faces[i+2]];

		uvIdx := rootIdx * 2
		tc1 := mgl32.Vec2{texCoords[uvIdx], texCoords[uvIdx+1]}
		tc2 := mgl32.Vec2{texCoords[uvIdx+2], texCoords[uvIdx+3]}
		tc3 := mgl32.Vec2{texCoords[uvIdx+4], texCoords[uvIdx+5]}
		//const vec2 &tc1 = texCoords[faces[i]];
		//const vec2 &tc2 = texCoords[faces[i+1]];
		//const vec2 &tc3 = texCoords[faces[i+2]];

		q1 := p2.Sub(p1)
		q2 := p3.Sub(p1)
		//vec3 q1 = p2 - p1;
		//vec3 q2 = p3 - p1;
		s1 := tc2.X() - tc1.X()
		s2 := tc3.X() - tc1.X()
		t1 := tc2.Y() - tc1.Y()
		t2 := tc3.Y() - tc1.Y()
		//float s1 = tc2.x - tc1.x, s2 = tc3.x - tc1.x;
		//float t1 = tc2.y - tc1.y, t2 = tc3.y - tc1.y;
		r := 1.0 / (s1*t2 - s2*t1)
		//float r = 1.0f / (s1 * t2 - s2 * t1);
		tan1 := mgl32.Vec3{
			(t2*q1.X() - t1*q2.X()) * r,
			(t2*q1.Y() - t1*q2.Y()) * r,
			(t2*q1.Z() - t1*q2.Z()) * r,
		}
		//vec3 tan1( (t2*q1.x - t1*q2.x) * r,
		//(t2*q1.y - t1*q2.y) * r,
		//(t2*q1.z - t1*q2.z) * r);

		tan2 := mgl32.Vec3{
			(s1*q2.X() - s2*q1.X()) * r,
			(s1*q2.Y() - s2*q1.Y()) * r,
			(s1*q2.Z() - s2*q1.Z()) * r,
		}
		//vec3 tan2( (s1*q2.x - s2*q1.x) * r,
		//(s1*q2.y - s2*q1.y) * r,
		//(s1*q2.z - s2*q1.z) * r);
		tan1Accum[i] += tan1.X()
		tan1Accum[i+1] += tan1.Y()
		tan1Accum[i+2] += tan1.Z()
		tan2Accum[i] += tan2.X()
		tan2Accum[i+1] += tan2.Y()
		tan2Accum[i+2] += tan2.Z()
		//tan1Accum[faces[i]] += tan1;
		//tan1Accum[faces[i+1]] += tan1;
		//tan1Accum[faces[i+2]] += tan1;
		//tan2Accum[faces[i]] += tan2;
		//tan2Accum[faces[i+1]] += tan2;
		//tan2Accum[faces[i+2]] += tan2;
	}

	for i := uint(0); i < uint(len(points))-2; i++ {
		n := mgl32.Vec3{
			normals[i],
			normals[i+1],
			normals[i+2],
		}
		t1 := mgl32.Vec3{
			tan1Accum[i],
			tan1Accum[i+1],
			tan1Accum[i+2],
		}
		t2 := mgl32.Vec3{
			tan2Accum[i],
			tan2Accum[i+1],
			tan2Accum[i+2],
		}
		//const vec3 &n = normals[i];
		//vec3 &t1 = tan1Accum[i];
		//vec3 &t2 = tan2Accum[i];

		// Gram-Schmidt orthogonalize
		//tangents[i] = vec4(glm::normalize( t1 - (glm::dot(n,t1) * n) ), 0.0f);
		res := t1.Sub(n.Mul(n.Dot(t1))).Normalize()
		tangents[i] = res.X()
		tangents[i+1] = res.Y()
		tangents[i+2] = res.Z()
		// Store handedness in w
		w := float32(1.0)
		if n.Cross(t1).Dot(t2) < 0 {
			w = -1.0
		}
		tangents[i+3] = w
		//tangents[i] = (glm::dot( glm::cross(n,t1), t2 ) < 0.0f) ? -1.0f : 1.0f;
	}

	//tan1Accum.clear();
	//tan2Accum.clear();

	return tangents
}
