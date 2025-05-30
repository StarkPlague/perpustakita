function fetchAndRenderBooks() {
                fetch('/books')
                .then(response => response.json())
                .then(data => {
                    const tableBody = document.getElementById("book-body")
                    tableBody.innerHTML = ""
                    data.forEach(book => {
                        const row = document.createElement("tr");
                        row.innerHTML = `
                            <td>${book.title}</td>
                            <td>${book.author}</td>
                            <td>${book.quantity}</td>
                            <td><button onclick="deleteBook(${book.id})">Hapus</button></td>
                            <td><button onclick="editBook(${book.id}, '${book.title}','${book.author}',${book.quantity})">Edit</button></td>
                           ` // tambahin ' ' pada input yang mengandung spasi
                        tableBody.appendChild(row);
                    });
                })
                .catch(err => console.error("Error fetching books: ", err))
            }

window.onload = function () {
    fetchAndRenderBooks();
};