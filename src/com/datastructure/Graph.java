/**
 * 
 */
package com.datastructure;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;

/**
 * @author Chris
 *
 */
public class Graph {
    private Map<Integer, ArrayList<Integer>> mAdjacencyList = new HashMap<>();

    public void addEdge(Integer mNode1, Integer mNode2) {
        ArrayList<Integer> mAdjacent = mAdjacencyList.get(mNode1);
        if (mAdjacent == null) {
            mAdjacent = new ArrayList<Integer>();
            mAdjacencyList.put(mNode1, mAdjacent);
        }
        mAdjacent.add(mNode2);
    }

    public void addTwoDirectionVertex(Integer mNode1, Integer mNode2) {
        addEdge(mNode1, mNode2);
        addEdge(mNode2, mNode1);
    }

    public boolean bIsConnected(Integer mNode1, Integer mNode2) {
        ArrayList<Integer> mList = mAdjacencyList.get(mNode1);
        if (mList == null) {
            return false;
        }
        return mList.contains(mNode2);
    }

    public ArrayList<Integer> adjecentNodes(Integer mNode) {
        ArrayList<Integer> mList = mAdjacencyList.get(mNode);
        if (mList == null) {
            return new ArrayList<Integer>();
        }
        return mList;
    }

    public void printAdjacencyList() {
        Iterator mIter = this.mAdjacencyList.entrySet().iterator();
        while (mIter.hasNext()) {
            Map.Entry mEntry = (Map.Entry) mIter.next();
            System.out.print(mEntry.getKey() + " : ");
            Iterator mIterList = ((ArrayList<Integer>) mEntry.getValue()).iterator();
            while (mIterList.hasNext()) {
                System.out.print(mIterList.next() + ", ");
            }
            System.out.println();
        }
    }

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        Graph mGraph = new Graph();
        mGraph.addEdge(1, 2);
        mGraph.addEdge(1, 3);
        mGraph.addEdge(2, 1);
        mGraph.addEdge(2, 4);
        mGraph.addEdge(2, 5);
        mGraph.addEdge(2, 6);
        mGraph.addEdge(3, 1);
        mGraph.addEdge(3, 2);
        mGraph.addEdge(3, 6);
        mGraph.addEdge(4, 2);
        mGraph.addEdge(5, 3);
        mGraph.addEdge(5, 6);
        mGraph.addEdge(6, 2);
        mGraph.addEdge(6, 3);
        mGraph.addEdge(6, 5);

        mGraph.addTwoDirectionVertex(5, 7);

        mGraph.printAdjacencyList();
    }

}
