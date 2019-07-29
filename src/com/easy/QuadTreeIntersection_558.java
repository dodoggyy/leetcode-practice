/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class QuadTreeIntersection_558 {
    // Definition for a QuadTree node.
    class Node {
        public boolean val;
        public boolean isLeaf;
        public Node topLeft;
        public Node topRight;
        public Node bottomLeft;
        public Node bottomRight;

        public Node() {
        }

        public Node(boolean _val, boolean _isLeaf, Node _topLeft, Node _topRight, Node _bottomLeft, Node _bottomRight) {
            val = _val;
            isLeaf = _isLeaf;
            topLeft = _topLeft;
            topRight = _topRight;
            bottomLeft = _bottomLeft;
            bottomRight = _bottomRight;
        }
    };

    // 2019年7月29日
    // Use recursive
    // Time Complexity: O(h)
    // Space Complexity:O(h)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 40 MB, less than 100.00%
    public Node intersect(Node quadTree1, Node quadTree2) {
        if (quadTree1.isLeaf) {
            return quadTree1.val ? quadTree1 : quadTree2;
        }
        if (quadTree2.isLeaf) {
            return quadTree2.val ? quadTree2 : quadTree1;
        }

        Node mTopLeft = intersect(quadTree1.topLeft, quadTree2.topLeft);
        Node mTopRight = intersect(quadTree1.topRight, quadTree2.topRight);
        Node mBottomLeft = intersect(quadTree1.bottomLeft, quadTree2.bottomLeft);
        Node mBottomRight = intersect(quadTree1.bottomRight, quadTree2.bottomRight);

        if (mTopLeft.val == mTopRight.val && mTopLeft.val == mBottomLeft.val && mTopLeft.val == mBottomRight.val
                && mTopLeft.isLeaf && mTopRight.isLeaf && mBottomLeft.isLeaf && mBottomRight.isLeaf) {
            return new Node(mTopLeft.val, true, null, null, null, null);
        } else {
            return new Node(false, false, mTopLeft, mTopRight, mBottomLeft, mBottomRight);
        }
    }
}
