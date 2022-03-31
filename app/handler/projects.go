package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/mingrammer/go-todo-rest-api-example/app/model"
	"net/http"
)

func GetAllProjects(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	//add := func(a, b int) int {
	//	return a + b
	//}
	//fmt.Println(add(3, 4))
	//fmt.Println(add1(3, 4))
	//fmt.Println(scope())
	nums := []int{10, 20, 30}

	fmt.Println(adder(nums...))

	//projects := []model.Project{}
	//
	//db.Where("created_at >=  ?", "2022-01-31").Find(&projects)
	//
	//var projectsResponseDTOs [3]model.ProjectResponseDTO
	//
	//for i, p := range projects {
	//	var projectResponseDTO model.ProjectResponseDTO
	//	projectResponseDTO.ID = p.ID
	//	projectResponseDTO.TitleResponse = p.Title + "Response1"
	//	projectResponseDTO.ArchivedResponse = !p.Archived
	//	projectsResponseDTOs[i] = projectResponseDTO
	//}
	//
	//fmt.Println(projectsResponseDTOs)
	//respondJSON(w, http.StatusOK, projects)
}

// Using ... before the type name of the last parameter indicates
// that it takes zero or more of those parameters.
// The function is invoked like any other function except we can
// pass as many arguments as we want.
func adder(args ...int) int {
	total := 0
	for i, v := range args { // Iterate over all args
		println(i)
		total += v
	}
	return total
}

func scope() func() int {
	outerVar := 2
	fmt.Println("outerVar=", outerVar)
	foo := func() int {
		fmt.Println("outerVar1=", outerVar)
		return outerVar
	}
	fmt.Println("outerVar2=", foo())

	return foo
}

func add1(a, b int) int {
	return a + b
}

func CreateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	project := model.Project{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

func UploadProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	//f, err := excelize.OpenFile("simple.xlsx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//c1, err := f.GetCellValue("Sheet1", "A1")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(f, err, c1)

	// Create a new sheet.
	//index := f.NewSheet("Sheet2")
	// Set value of a cell.
	//f.SetCellValue("Sheet2", "A2", "Hello world.")
	//f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	//f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	//if err := f.SaveAs("Book1.xlsx"); err != nil {
	//	println(err.Error())
	//}

	//f, err := excelize.OpenFile("simple.xlsx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//c1, err := f.GetCellValue("Sheet1", "A1")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(f, err)
	//
	//r.ParseMultipartForm(10 << 20)
	//file, handler, err := r.FormFile("myFile")
	//
	//if err != nil {
	//	fmt.Println("Error Retrieving the File")
	//	fmt.Println(err)
	//	return
	//}
	//
	//f, err := excelize.OpenFile("simple.xlsx")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//c1, err := f.GetCellValue("Sheet1", "A1")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(c1)
	//
	//fileBytes, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("read file excel: ", fileBytes)
	//
	//defer file.Close()

}

func GetProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db, title, w, r)
	if project == nil {
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func UpdateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db, title, w, r)
	if project == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func DeleteProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db, title, w, r)
	if project == nil {
		return
	}
	if err := db.Delete(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func ArchiveProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db, title, w, r)
	if project == nil {
		return
	}
	project.Archive()
	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func RestoreProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db, title, w, r)
	if project == nil {
		return
	}
	project.Restore()
	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// getProjectOr404 gets a project instance if exists, or respond the 404 error otherwise
func getProjectOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *model.Project {
	project := model.Project{}
	if err := db.First(&project, model.Project{Title: title}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &project
}
