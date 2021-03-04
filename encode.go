package smd

import (
	"fmt"
	"io"
)

func (smd *SMD) Encode(w io.Writer) error {
	fmt.Fprintln(w, "version 1")
	if smd.Nodes != nil {
		fmt.Fprintln(w, "nodes")
		for _, node := range smd.Nodes {
			fmt.Fprintf(w, "  %v \"%v\" %v\n", node.ID, node.Name, node.ParentID)
		}
		fmt.Fprintln(w, "end")
	}
	if smd.Skeleton != nil {
		fmt.Fprintln(w, "skeleton")
		times := len(smd.Skeleton)
		for time := 0; time < times; time++ {
			bones := smd.Skeleton[time]
			fmt.Fprintf(w, "  time %v\n", time)
			for _, bone := range bones {
				fmt.Fprintf(
					w, "    %d %.6f %.6f %.6f %.6f %.6f %.6f\n", bone.BoneID,
					bone.Pos[0], bone.Pos[1], bone.Pos[2],
					bone.Rot[0], bone.Rot[1], bone.Rot[2],
				)
			}
		}
		fmt.Fprintln(w, "end")
	}
	if smd.Triangles != nil {
		fmt.Fprintln(w, "triangles")
		for _, tri := range smd.Triangles {
			fmt.Fprintln(w, tri.Material)
			for _, v := range tri.Vertexes {
				fmt.Fprintf(
					w, "  %d %.6f %.6f %.6f %.6f %.6f %.6f %.6f %.6f",
					v.ParentBoneID,
					v.Pos[0], v.Pos[1], v.Pos[2],
					v.Norm[0], v.Norm[1], v.Norm[2],
					v.UV[0], v.UV[1],
				)
				if v.Links != nil && len(v.Links) > 0 {
					fmt.Fprintf(w, " %d", len(v.Links))
					for _, link := range v.Links {
						fmt.Fprintf(w, " %d %.6f", link.BoneID, link.Weight)
					}
				}
				fmt.Fprintln(w)
			}
		}
		fmt.Fprintln(w, "end")
	}
	return nil
}
