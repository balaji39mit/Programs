package main

import "fmt"

func (e *Employee) Find(Id int64) *Employee {
	if e.Id == Id {
		return e
	}
	for _, val := range e.Children {
		if employee := val.Find(Id); employee != nil {
			return employee
		}
	}
	return nil
}

func (e *Employee) Print(depth int) {
	for i := len(e.Children) - 1; i > len(e.Children)/2-1; i-- {
		e.Children[i].Print(depth + 1)
	}
	for i := 0; i < depth; i++ {
		fmt.Print("\t\t\t\t\t\t")
	}
	fmt.Println(fmt.Sprintf("Name:%v (Position:%v)", e.Name, e.Position))
	fmt.Print("\n\n")
	for i := len(e.Children)/2 - 1; i >= 0; i-- {
		e.Children[i].Print(depth + 1)
	}
}

func (e *Employee) SplitBonus(bonus float64) {
	children := len(e.Children)
	if children <= 0 {
		return
	}
	e.Bonus = bonus
	averageAmount := bonus / float64(children)
	for _, val := range e.Children {
		val.Bonus = averageAmount * (BonusPercentage[val.Rating] / 100)
		if val.Children != nil {
			//This employee is the last node. We don't need to split further.
			val.SplitBonus(val.Bonus)
		}
	}
}

func (e *Employee) AnalyzePerformance(performers *map[float64]*Employee, number int, position Position) {
	if e.Position == position {
		fmt.Println("analysis..")
		ratio := e.Bonus / e.Salary
		//capacity is not reached, just append.
		if len(*performers) < number {
			(*performers)[ratio] = e
		} else if len(*performers) == number { //capacity has been reached, just compare it with lowest element.
			minimumIndex := MAXFLOAT
			for index := range *performers {
				if index < minimumIndex {
					minimumIndex = index
				}
			}
			delete(*performers, minimumIndex)
			(*performers)[ratio] = e
			return
		}
		return
	}
	for _, val := range e.Children {
		val.AnalyzePerformance(performers, number, position)
	}
}

func (e *Employee) LowestLevel() Position {
	if e.Children == nil || len(e.Children) <= 0 {
		return e.Position
	}
	//TODO:Enhance it to process the lowest level position for any given employee.
	return e.Position
}

/*Not needed.
func (e *Employee) AnalyzeDEPerformance(performer []*Employee, queue []float64) {
	if e.Position == DeliveryExecutives {
		ratio := e.Bonus / e.Salary
		//capacity is not reached, just append.
		if cap(queue) != len(queue) {
			queue = append(queue, ratio)
			performer = append(performer, e)
			sort.Float64s(queue)
		} else { //capacity has been reached, just compare it with lowest element.
			if cap(queue) > 0 {
				queue[0] = ratio
				//figure out how to update the performer.
			}
		}
		return
	}
	for _, val := range e.Children {
		val.AnalyzeDEPerformance(performer, queue)
	}
}
*/
