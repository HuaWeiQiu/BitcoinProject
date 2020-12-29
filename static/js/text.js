
function modifyPwdSumbit(){
    if($('#oldPassword').val()==""){
        alert('原始密码为空');
        return
    }
    if($('#newPassword').val()==""){
        alert('新密码为空');
        return
    }
    if($('#repeatPassword').val()==""){
        alert('确认密码为空');
        return
    }
    var data={
        'oldPassword':$('#oldPassword').val(),
        'newPassword':$('#newPassword').val(),
        'repeatPassword':$('#repeatPassword').val()
    }
    $.ajax({
        type:"POST",
        url:"user/modifyPwd",
        dataType:"json",
        contentType:"application/json",
        data:JSON.stringify(data),
        success:function(res){
            console.log(res)
            if(res.code == 200){
                alert(res.msg)
                showModifyPwdPage()
            }
        }
    });
}