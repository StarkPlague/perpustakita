function deleteBook(id) {
                    fetch('/delete-book', {
                        method: 'POST',
                        headers: {
                            'Content-Type' : 'application/x-www-form-urlencoded',
                        },
                        body: 'id=' + encodeURIComponent(id)
                    })
                    .then(res => {
                        if (res.ok) {
                            location.reload();
                        } else {
                            alert("Gagal menghapus buku")
                        }
                    })
                }

function editBook(id,title,author,quantity) {
    const newTitle = prompt("Edit Judul: ", title)
    const newAuthor = prompt("Edit Penulis: ", author)
    const newQuantity = prompt("Edit jumlah: ", quantity)

    if (newTitle && newAuthor && newQuantity) {
        fetch('/update-book', {
            method: 'POST',
            headers: {
                'Content-Type':'application/x-www-form-urlencoded'
            },
            body: `id=${id}&title=${encodeURIComponent(newTitle)}&author=${encodeURIComponent(newAuthor)}&quantity=${newQuantity}`
        }).then(res =>{
            if (res.ok) {
                location.reload()
            } else {
                alert("Gagal update buku")
            }
        })
    }
}

