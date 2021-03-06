package mysql

import (
	"errors"
    "database/sql"
    "time"

    // Import the models package that we just created. You need to prefix this with 
    // whatever module path you set up back in chapter 02.02 (Project Setup and Enabling 
    // Modules) so that the import statement looks like this:
    // "{your-module-path}/pkg/models".
    "github.com/cyril-ui-developer/codersnippet/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
    DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content string, expires time.Time ) (int, error) {
	
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
    
	result, err := m.DB.Exec(stmt, title, content, expires)
    if err != nil {
        return 0, err
    }
    
	id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }


    return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE id = ?`

	 // Use the QueryRow() method on the connection pool to execute our
    // SQL statement
    row := m.DB.QueryRow(stmt, id)

    // Initialize a pointer to a new zeroed Snippet struct.
    s := &models.Snippet{}
   
	// Use row.Scan() to copy the values from each field in sql.Row to the
    // corresponding field in the Snippet struct. Notice that the arguments
    // to row.Scan are *pointers* to the place you want to copy the data into,
    // and the number of arguments must be exactly the same as the number of
    // columns returned by your statement.
    err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err != nil {
        // If the query returns no rows, then row.Scan() will return a
        // sql.ErrNoRows error. We use the errors.Is() function check for that
        // error specifically, and return our own models.ErrNoRecord error
        // instead.
        if errors.Is(err, sql.ErrNoRows) {
            return nil, models.ErrNoRecord
        } else {
             return nil, err
        }
    }

    // If everything went OK then return the Snippet object.
    return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    ORDER BY created DESC LIMIT 100`
    
	 // Use the Query() method on the connection pool to execute our
    // SQL statement. 
	rows, err := m.DB.Query(stmt)
    if err != nil {
        return nil, err
    }
    
	defer rows.Close()
    
	// Initialize an empty slice to hold the models.Snippets objects.
	snippets := []*models.Snippet{}

	 // Use rows.Next to iterate through the rows in the resultset.
	 for rows.Next() {
        // Create a pointer to a new zeroed Snippet struct.
        s := &models.Snippet{}
        // Use rows.Scan() to copy the values from each field in the row to the
        // new Snippet object that we created. Again, the arguments to row.Scan()
        // must be pointers to the place you want to copy the data into, and the
        // number of arguments must be exactly the same as the number of
        // columns returned by your statement.
        err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
        if err != nil {
            return nil, err
        }
        // Append it to the slice of snippets.
        snippets = append(snippets, s)
    }

	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
    // error that was encountered during the iteration.

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return snippets, nil
}
