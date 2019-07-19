package com.easy;

import java.util.List;

public class Node {
    // 2019年7月19日
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val,List<Node> _children) {
        val = _val;
        children = _children;
    }
}
