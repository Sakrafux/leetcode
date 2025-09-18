package main

import "fmt"

func main() {
	tm := Constructor([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}})
	tm.Add(4, 104, 5)
	tm.Edit(102, 8)
	fmt.Println(tm.ExecTop()) // 3
	tm.Rmv(101)
	tm.Add(5, 105, 15)
	fmt.Println(tm.ExecTop()) // 5

	tm2 := Constructor([][]int{{1, 101, 8}, {2, 102, 20}, {3, 103, 5}})
	tm2.Add(4, 104, 5)
	tm2.Edit(102, 9)
	fmt.Println(tm2.ExecTop()) // 2
	tm2.Rmv(101)
	tm2.Add(50, 101, 8)
	fmt.Println(tm2.ExecTop()) // 50

	tm3 := Constructor([][]int{{10, 26, 25}})
	tm3.Rmv(26)
	fmt.Println(tm3.ExecTop()) // -1

	tm4 := Constructor([][]int{{1, 30, 22}, {0, 11, 27}, {9, 15, 30}, {4, 17, 37}, {6, 16, 12}, {4, 24, 35}, {0, 13, 35}, {0, 14, 2}, {3, 10, 34}})
	fmt.Println(tm4.ExecTop()) // 4
	tm4.Edit(24, 30)
	tm4.Rmv(14)
	fmt.Println(tm4.ExecTop()) // 0

	tm5 := Constructor([][]int{
		{20, 55, 91}, {30, 94, 511}, {35, 90, 877}, {40, 57, 281}, {33, 11, 408},
		{13, 91, 217}, {31, 95, 239}, {8, 51, 132}, {42, 77, 30}, {20, 36, 894},
		{12, 15, 43}, {1, 48, 591}, {22, 85, 420}, {21, 83, 767}, {6, 25, 894},
		{38, 38, 722}, {42, 65, 579}, {47, 54, 436}, {37, 45, 483}, {27, 74, 712},
		{23, 63, 610}, {48, 24, 843}, {26, 81, 742}, {41, 78, 625}, {27, 75, 506},
		{13, 70, 121}, {24, 6, 103},
	})

	tm5.Add(12, 71, 260)
	tm5.Add(26, 30, 626)
	tm5.Add(1, 68, 206)
	tm5.Edit(55, 566)
	tm5.Edit(6, 312)
	tm5.Edit(36, 919)
	tm5.Edit(95, 720)
	tm5.Rmv(55)
	tm5.Add(9, 29, 635)
	tm5.Edit(54, 450)
	fmt.Println(tm5.ExecTop())
	tm5.Rmv(90)
	tm5.Add(28, 82, 378)
	tm5.Add(21, 66, 569)
	tm5.Edit(30, 476)
	tm5.Edit(48, 500)
	tm5.Rmv(85)
	tm5.Edit(6, 246)
	tm5.Rmv(71)
	fmt.Println(tm5.ExecTop())
	tm5.Add(40, 37, 335)
	fmt.Println(tm5.ExecTop())
	tm5.Rmv(95)
	fmt.Println(tm5.ExecTop())
	tm5.Add(21, 67, 174)
	fmt.Println(tm5.ExecTop())
	fmt.Println(tm5.ExecTop())
	tm5.Rmv(74)
	tm5.Rmv(63)
	fmt.Println(tm5.ExecTop())
	fmt.Println(tm5.ExecTop())
	fmt.Println(tm5.ExecTop())
	fmt.Println(tm5.ExecTop())
}

// Solution

// A helper struct for the heap elements
type Task struct {
	userId   int
	taskId   int
	priority int
	heapPos  int // It needs to keep track of its position in order to remove it at will
}

type TaskHeap []*Task

func NewTaskHeap() TaskHeap {
	// 1-indexed heap is just easier to understand, but otherwise doesn't matter
	return make(TaskHeap, 1)
}

func (h TaskHeap) len() int {
	return len(h)
}

// Append new task to the end of the heap and heapify up
func (h TaskHeap) push(task *Task) TaskHeap {
	h = append(h, task)
	task.heapPos = h.len() - 1
	h.up(h.len() - 1)
	return h
}

// Heapify up
func (h TaskHeap) up(index int) {
	if index == 1 {
		return
	}

	parentIndex := index / 2

	task := h[index]
	parentTask := h[parentIndex]

	if Less(task, parentTask) {
		// Switch places
		h[index], h[parentIndex] = parentTask, task
		// Don't forget to swap position information as well
		task.heapPos, parentTask.heapPos = parentTask.heapPos, task.heapPos
		// Recursively heapify
		h.up(parentIndex)
	}
}

// Heapify down
func (h TaskHeap) down(index int) {
	if index*2 >= h.len() {
		return
	}

	task := h[index]

	childIndex1 := index * 2
	childTask1 := h[childIndex1]

	if index*2+1 >= h.len() {
		if Less(childTask1, task) {
			h[index], h[childIndex1] = childTask1, task
			task.heapPos, childTask1.heapPos = childTask1.heapPos, task.heapPos
		}
		return
	}

	childIndex2 := index*2 + 1
	childTask2 := h[childIndex2]

	if Less(childTask1, childTask2) {
		if Less(childTask1, task) {
			h[index], h[childIndex1] = childTask1, task
			task.heapPos, childTask1.heapPos = childTask1.heapPos, task.heapPos
			h.down(childIndex1)
		}
	} else {
		if Less(childTask2, task) {
			h[index], h[childIndex2] = childTask2, task
			task.heapPos, childTask2.heapPos = childTask2.heapPos, task.heapPos
			h.down(childIndex2)
		}
	}
}

// Remove a task at will from the heap
func (h TaskHeap) remove(index int) TaskHeap {
	task := h[index]
	lastTask := h[h.len()-1]

	// Move the last task at the location of the current task
	h[index], h[h.len()-1] = h[h.len()-1], lastTask

	// And shrink the heap
	h = h[:h.len()-1]
	lastTask.heapPos = task.heapPos
	// Heapify the swapped element down...
	h.down(index)
	// ...or try to heapify it up
	if index < h.len() {
		h.up(index)
	}

	return h
}

func Less(a, b *Task) bool {
	return a.priority > b.priority || (a.priority == b.priority && a.taskId > b.taskId)
}

type TaskManager struct {
	taskMap  map[int]*Task
	taskHeap TaskHeap
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{make(map[int]*Task), NewTaskHeap()}

	for _, task := range tasks {
		t := &Task{task[0], task[1], task[2], 0}
		tm.taskMap[t.taskId] = t
		tm.taskHeap = tm.taskHeap.push(t)
	}

	return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	t := &Task{userId, taskId, priority, 0}
	this.taskMap[t.taskId] = t
	// Simply add task to heap
	this.taskHeap = this.taskHeap.push(t)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	t := this.taskMap[taskId]
	t.priority = newPriority
	// Remove the task from the heap and push it anew to guarantee correct ordering
	this.taskHeap = this.taskHeap.remove(t.heapPos).push(t)
}

func (this *TaskManager) Rmv(taskId int) {
	t := this.taskMap[taskId]
	// Simply remove from heap
	this.taskHeap = this.taskHeap.remove(t.heapPos)
}

func (this *TaskManager) ExecTop() int {
	if this.taskHeap.len() == 1 {
		return -1
	}

	t := this.taskHeap[1]
	// Simply remove from heap
	this.taskHeap = this.taskHeap.remove(t.heapPos)
	return t.userId
}
