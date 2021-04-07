new Vue({

    el: "#app",

    data: {
        students: [],

        stuname: "",
        info: "",
        current: -1,
        options: [
            { value: -1, label: "状态无" },
            { value: 0, label: "状态No" },
            { value: 1, label: "状态Yes" },
        ],
        isEntered: false
    },

    computed: {
        labels() {
            return this.options.reduce(function(a, b) {
                return Object.assign(a, {[b.value]: b.label})
            }, {})
        },

        computedStudents() {
            return this.students.filter(function(el) {
                var option = this.current < 0 ? true : this.current === el.state
                return option
                
            }, this)

        },

        // 验证输入的名字
        validate() {
            var isEnteredStudentName = 0 < this.stuname.length
            this.isEntered = isEnteredStudentName
            return isEnteredStudentName
        }
    },

    created: function() {
        this.getAllStudents()
    },

    methods: {
        getAllStudents() {
            axios.get("/v1/student/getAll")
            .then(response => {
                if (response.status != 200) {
                    throw new Error("返回数据错误")
                } else {
                    var result = response.data
                    this.students = result
                }
            })
        },

        getStudent(student) {
            axios.get("/v1/student/get", {
                params: {
                    id: student.id
                }
            })
            .then(response => {
                if (response.status != 200) {
                    throw new Error("获取信息错误")
                } else {
                    var res = response.data
                    console.log(res)
                    var index = this.students.indexOf(student)
                    this.students.splice(index, 1, res)
                }
            })
        },

        addStudent() {
            const params = new URLSearchParams();
            params.append("name", this.stuname)
            params.append("info", this.info)

            axios.post("/v1/student/add", params)
            .then(response => {
                if (response.status != 200) {
                    throw new Error("增加错误")
                } else {
                    this.getAllStudents()

                    this.initInputValue()
                }
            })
        },

        changeStudentState(student) {
            const params = new URLSearchParams();
            params.append("id", student.id)
            params.append("status", student.status)

            axios.post("/v1/student/changestatus", params)
            .then(response => {
                if (response.status != 200) {
                    throw new Error("修改状态错误")
                } else {
                    this.getStudent(student)
                }
            })
        },

        deleteStudent(student) {
            const params = new URLSearchParams();
            params.append("id", student.id)

            axios.post("/v1/student/delete", params)
            .then(response => {
                if (response.status != 200) {
                    throw new Error("删除错误")
                } else {
                    this.getAllStudents()
                }
            })
        },

        // 初始化
        initInputValue() {
            this.current = -1
            this.stuname = ""
            this.info = ""
        }
    }
})