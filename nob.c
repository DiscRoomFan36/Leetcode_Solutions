
#define NOB_IMPLEMENTATION
#define NOB_STRIP_PREFIX
#include "thirdparty/nob.h"

#define DAYS_FOLDER         "./days/"
#define BUILD_FOLDER        "./build/"
#define THIRDPARTY_FOLDER   "./thirdparty/"



static Cmd cmd = {0};


int main(int argc, char **argv) {
    NOB_GO_REBUILD_URSELF(argc, argv);

    mkdir_if_not_exists(BUILD_FOLDER);


    // get a list of all days
    File_Paths day_folders = {0};
    if (!read_entire_dir(DAYS_FOLDER, &day_folders))  exit(EXIT_FAILURE);

    // remove all things that start with '.'
    for (size_t i = 0; i < day_folders.count; i++) {
        const char *folder = day_folders.items[i];
        assert(folder);
        if (folder[0] == '.') {
            da_remove_unordered(&day_folders, i);
            i -= 1;
        }
    }

    // sort by name
    qsort(day_folders.items, day_folders.count, sizeof(*day_folders.items), (__compar_fn_t) strcmp);

    // compile all days
    for (int i = 0; i < day_folders.count; i++) {
        size_t checkpoint = temp_save();

        const char *folder = day_folders.items[i];
        printf("%d: %s\n", i, folder);
        const char *folder_name = temp_sprintf(DAYS_FOLDER"%s", folder);
        assert(get_file_type(folder_name) == NOB_FILE_DIRECTORY);

        File_Paths files = {};
        read_entire_dir(folder_name, &files);


        // get the index of the main file.

        #define MAIN_DOT "main."
        int main_index = -1;
        for (int j = 0; j < day_folders.count; j++) {
            String_View sv = sv_from_cstr(files.items[j]);
            if (sv_starts_with(sv, sv_from_cstr(MAIN_DOT))) {
                assert(main_index == -1 && "two files cannot both start with '"MAIN_DOT"'");
                main_index = j;
            }
        }
        assert(main_index != -1 && "must contain a file start starts with '"MAIN_DOT"'");

        String_View main_filename = sv_from_cstr(files.items[main_index]);
        if (sv_end_with(main_filename, ".cpp")) {
            mkdir_if_not_exists(temp_sprintf(BUILD_FOLDER"%s", folder));

            // compile a c++ file
            cmd_append(&cmd, "clang++");

            cmd_append(&cmd, "-o", temp_sprintf(BUILD_FOLDER"%s/main", folder));
            cmd_append(&cmd, temp_sprintf("%s/"SV_Fmt, folder_name, SV_Arg(main_filename)));

            cmd_append(&cmd, "-Wall", "-Wextra");

            if (!cmd_run(&cmd)) exit(EXIT_FAILURE);

        } else {
            nob_log(ERROR, "unknown file extention on file '"SV_Fmt"'", SV_Arg(main_filename));
        }


        free(files.items);
        temp_rewind(checkpoint);
    }


    exit(EXIT_SUCCESS);
}

