package 27-remove-element

func removeElement(nums []int, val int) int {
    outputIndex := 0
    for i, _ := range nums {
        elem := nums[i]
        if elem != val {
            nums[outputIndex] = elem
            outputIndex += 1
        } 
    }    
    return outputIndex
}