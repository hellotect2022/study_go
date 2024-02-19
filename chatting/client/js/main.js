async function fetchHtmlAsText(url) {
    return await (await fetch(url)).text();
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
    //fetchData();
}

function renderMainComponents(){
    var buttons = document.getElementsByClassName("ui-buttons");
    Array.from(buttons).forEach(function(button) {
        button.addEventListener("click", function(event) {
            // 1. content-* 탭 모두 hidden 처리
            var contentElements = document.querySelectorAll('[id^="content-"]');
            contentElements.forEach(element => element.classList.add('hidden'))
            // 2. 선택된 버튼과 관련된 탭만 show
            var selectedContent = event.target.getAttribute('name')
            document.getElementById(`content-${selectedContent}`).classList.remove('hidden')
        });
    });
}
