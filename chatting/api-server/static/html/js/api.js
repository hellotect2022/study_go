async function loginApi(event) {
    console.log('???')
    event.preventDefault(); // 기본 제출 동작을 막습니다.

        // 폼 데이터를 수집하여 JavaScript 객체로 변환합니다.
    const formData = {
        username: encodeURIComponent(document.querySelector("#username").value),
        password: encodeURIComponent(document.querySelector("#password").value)
    };

    try {
        // AJAX 요청을 보냅니다.
        const response = await fetch("/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(formData)
        });

        // 응답을 JSON으로 변환합니다.
        const data = await response.json();

        console.log('result', data);

        // 처리된 값에 따라 다른 경로로 라우팅합니다.
        if (data.status === "success") {
            // 로컬 스토리지에 데이터 저장
            localStorage.setItem('userId', formData.username);
            window.location.href = "/view/chatRoom.html"; // 성공했을 때의 경로
        } 

    } catch (error) {
        console.error("Error:", error);
        // 오류가 발생했을 때의 처리를 추가할 수 있습니다.
    }
}

async function getUserListAll(event) {
    event.preventDefault(); // 기본 제출 동작을 막습니다.
    try {
        // AJAX 요청을 보냅니다.
        const response = await fetch("/getUserListAll", {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            }
        });

        // 응답을 JSON으로 변환합니다.
        const data = await response.json();

        console.log('result', data);

        // 처리된 값에 따라 다른 경로로 라우팅합니다.
        if (data.status === "success") {
            console.log(data.result)
            for (const key in data.result) {
                // 로컬 스토리지에 데이터 저장
                var targetNode = document.getElementById("user-list");
                var newDiv = document.createElement("div");
                newDiv.textContent=key;
                newDiv.setAttribute('id',key)
                if (key != localStorage.getItem('userId')){
                    newDiv.addEventListener('click',function(event){
                        console.log('저는 이것 입니다.', event.target)
                        if (confirm(`${event.target.textContent} 와 채팅을 하시겠습니까?`)) {
                            // 방을 생성한다. 
                            var userMember = event.target.getAttribute('id')
                            createRoom(userMember);
                        }
                    })
                }
                
                targetNode.appendChild(newDiv);
            }
            //return data.result
        } 

    } catch (error) {
        console.error("Error:", error);
        // 오류가 발생했을 때의 처리를 추가할 수 있습니다.
    }
}

async function createRoom(joinUserId){
    try {
        const formData = {};
        var creatorId = localStorage.getItem('userId');
        formData.creator = creatorId;
        formData.members = [joinUserId, creatorId]
        // AJAX 요청을 보냅니다.
        const response = await fetch("/createRoom", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(formData)
        });

        // 응답을 JSON으로 변환합니다.
        const data = await response.json();

        // 처리된 값에 따라 다른 경로로 라우팅합니다.
        if (data.status === "success") {
            console.log("createRoom-message : ", data)
            createRoomUIComponent(data.result);
        }
    }catch (error) {
        console.error("Error:", error);
        // 오류가 발생했을 때의 처리를 추가할 수 있습니다.
    }
}

function createRoomUIComponent(params){
    var targetNode = document.getElementById("room-list");
    var newDiv = document.createElement("div");
    newDiv.textContent=params.members;
    newDiv.setAttribute('id',params.roomId)
    newDiv.addEventListener('click', event=>{
        console.log("소켓연결을 해야함",params)
        const joinRoomMessage = {
            type: 'joinRoom',
            userId: params.creator,
            message: params.members
        };
        socket.send(JSON.stringify(joinRoomMessage)); // JSON 형태로 요청 전송
    })
    targetNode.appendChild(newDiv);
}



/*
    //MutationObserver: MutationObserver를 사용하여 DOM 요소의 변화를 관찰하고 이에 대응하는 콜백 함수를 호출할 수 있습
    var observer = new MutationObserver(function(mutations) {
        console.log('mutations?? :', mutations);
        mutations.forEach(function(mutation){
            console.log('개별 mutation : ', mutation)
        })
    })
    
    var targetNode = document.getElementById("user-list");
    var config = { attributes: true, childList: true, subtree: true };

    observer.observe(targetNode,config);

    // document.querySelector("#login-form").addEventListener("submit", loginApi(event));
    */