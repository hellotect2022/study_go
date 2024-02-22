async function fetchHtmlAsText(url) {
    return (await fetch(url)).text();
}
async function importPage(target, callback) {
    var containers = document.querySelectorAll(`[class^="container-"]`);
    Array.from(containers).forEach(container => container.classList.add('hidden'));
    document.querySelector('.' + target).innerHTML = await fetchHtmlAsText(target + '.html');
    document.querySelector('.' + target).classList.remove('hidden');

    // 콜백 함수가 전달되었는지 확인 후 실행
    if (typeof callback === 'function') {
        await callback();
    } else {
        console.error('callback is not a function');
    }
}

function initMain(){
    //1. 페이지에 관련된 UI components 및 이벤트 할당
    renderMainComponents();
    //2. 페이지에 필요한 데이타를 받아오는 기능구현
    fetchData();
}

function renderMainComponents(){
    document.querySelectorAll('[id^="content-"]').forEach(el => {
        el.addEventListener('change',e=>console.log('change2->',e))
    })

    console.log('??',document.getElementsByClassName("buttons")[0].children)
    var buttons = document.getElementsByClassName("buttons")[0].children;
    Array.from(buttons).forEach(function(button) {
        button.addEventListener("click", function(event) {
            //1 모든 버튼 선택 CSS 초기화 
            Array.from(buttons).forEach(button=>button.classList.remove('selection'))
            // 3. content-* 탭 모두 hidden 처리
            var contentElements = document.querySelectorAll('[id^="content-"]');
            contentElements.forEach(element => {
                //element.classList.add('hidden');
                element.style.zIndex=0;
            })

            actionSelected(event.target);
        });
    });
}

// 버튼 선택시 해당 page 에 맞는 이벤트 발생
function actionSelected(node) {
    // 1 선택한 버튼 색상을 수정
    node.classList.add('selection')
    // 2. 선택된 버튼과 관련된 탭만 show
    var selBtnNm = node.getAttribute('name')
    //document.getElementById(`content-${selBtnNm}`).classList.remove('hidden')
    switch (selBtnNm){
        case "chattingRoom":
            openChattingRoomTest();
            break;
        case "chattingList":
            document.getElementById("content-"+selBtnNm).style.zIndex=3;
            break;
        case "userList":
            document.getElementById("content-"+selBtnNm).style.zIndex=3;
        break;
    }
}

async function openChattingRoomTest() {
    document.querySelector(".content").innerHTML += await fetch('chat-room.html')
        .then(response=>response.text())
        .then(data=>{
            return data
        })
    var room = document.getElementById("content-chattingRoom");
}

function fetchData(){
    var test = [0,1,2]
        //,3,4,5,6,7,8,9,0,1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6,7,8,9];
    drawUserList(test);
}

async function drawUserList(test) {
    var contentUserList = document.getElementById("content-userList").children[0];
    for (var i=0; i<test.length;i++){
        var li = document.createElement("li");
        li.innerHTML = await fetchHtmlAsText('user-profile.html');
        li.querySelector('.user-profile').id = test[i];
        li.querySelector('.user-bio>span').textContent = test[i];
        li.querySelector('.user-bio>div').textContent=i+" 입니다ㄹㄴㅇㄹㄹㄹㄹㄹㄹㄹdddddddddddddddddddddddd";
        if (test[i]%2 == 0){
            li.getElementsByClassName('user-image')[0].classList.add('user-image-girl');
        }else {
            li.getElementsByClassName('user-image')[0].classList.add('user-image-man');
        }
        contentUserList.appendChild(li)
    }    
} 


// 사용자 선태시 해당 사용자의 방으로 넘어감 
function letsTalkWithPeople(e){
    var user = {};
    user.id=e.getAttribute('id');
    user.name=e.children[1].children[0].textContent;

    // document.getElementById(`content-chattingRoom`).classList.remove('hidden')

    openChattingRoom(user);
}

async function openChattingRoom(user){
    //var currentOpendRoom = document.getElementsByClassName
    const elements = document.querySelectorAll('[id^="content-chatting-Room-"]');
    if (document.getElementById("content-chattingRoom-"+user.id) != null) {
        console.log("현재 생성된 방이 존재");
    }else {
        console.log("새로운 방을 생성 ->", user.id);
        document.querySelector(".content").innerHTML += await fetch('chat-room.html')
        .then(response=>response.text())
        .then(data=>{
            return data.replace("content-chattingRoom","content-chattingRoom-"+user.id)
        })
    }
    var room = document.getElementById("content-chattingRoom-"+user.id);
    // targetUser 프로필
    //room.children[0].children[1].setAttribute('src',"../resources/user_man.svg")
    // targetUser 이름
    room.children[0].children[2].textContent = user.name;
    // room 의 zIndex
    room.style.zIndex=4;
}

function goBack(e){
    var parentNode = e.parentNode;
    parentNode.style.zIndex=0;
}