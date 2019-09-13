package majority_element

import "errors"

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	s := NewStack()

	for _, v := range nums {

		if s.empty() {
			s.push(v)
			continue
		}

		if s.top() == v {
			s.push(v)
			continue
		}

		s.pop()
	}

	return s.top()
}

type stack struct {
	data []int
}

func NewStack() *stack {
	return &stack{
		data: make([]int, 0, 4),
	}
}

func (s *stack) push(v int) {
	s.data = append(s.data, v)

}

func (s *stack) len() int {
	return len(s.data)
}

func (s *stack) empty() bool {
	return s.len() == 0
}

func (s *stack) top() int {
	return s.data[len(s.data)-1]
}

func (s *stack) pop() (int, error) {
	if s.len() == 0 {
		return 0, errors.New("stack is empty")
	}

	v := s.top()
	s.data = s.data[:len(s.data)-1]

	return v, nil
}
