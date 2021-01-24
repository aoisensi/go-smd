package smd

type Node struct {
	ID       int
	Name     string
	ParentID int
}

type Skeleton map[int][]*SkeletonBone

type SkeletonBone struct {
	BoneID int
	Pos    [3]float64
	Rot    [3]float64
}

type Triangle struct {
	Material string
	Vertexes [3]*TriangleVertex
}

type TriangleVertex struct {
	ParentBoneID int
	Pos          [3]float64
	Norm         [3]float64
	UV           [2]float64
	Links        []*TriangleVertexLink
}

type TriangleVertexLink struct {
	BoneID int
	Weight float64
}
