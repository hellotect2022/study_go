function loginApi(event) {
    console.log('???')
    event.preventDefault(); // 기본 제출 동작을 막습니다.

        // 폼 데이터를 수집하여 JavaScript 객체로 변환합니다.
    const formData = {
        username: encodeURIComponent(document.querySelector("#username").value),
        password: encodeURIComponent(document.querySelector("#password").value)
    };

    // AJAX 요청을 보냅니다.
    fetch("/postLogin", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
    .then(response => response.json()) // 응답을 JSON으로 변환합니다.
    .then(data => {
        console.log('result',data)
        // 처리된 값에 따라 다른 경로로 라우팅합니다.
            if (data.status == "success") {
                window.location.href = "/chatRoom.html"; // 성공했을 때의 경로
            } 
    })
    .catch(error => {
        console.error("Error:", error);
        // 오류가 발생했을 때의 처리를 추가할 수 있습니다.
    });
}

function connectSocket(){
    return new WebSocket('ws://localhost:7777/ws');
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